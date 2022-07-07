package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func SplitHandler(w http.ResponseWriter, r *http.Request) {
	var (
		request  Request
		response Response
	)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(request.SplitInfo) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	splitInfo := request.SplitInfo
	SortSplitInfoSlice(splitInfo)

	if len(splitInfo) > 20 {
		splitInfo = splitInfo[:20]
	}
	response.SplitBreakDown = make([]SplitBreakDown, len(splitInfo))
	tracker := new(BalanceTracker)
	tracker.Balance = request.Amount
	for index, info := range splitInfo {
		switch info.SplitType {
		case "FLAT":
			tracker.Balance -= info.SplitValue
			response.SplitBreakDown[index] = SplitBreakDown{
				SplitEntityID: info.SplitEntityID,
				Amount:        info.SplitValue,
			}
		case "PERCENTAGE":
			amount := (info.SplitValue / 100) * tracker.Balance
			tracker.Balance -= amount
			response.SplitBreakDown[index] = SplitBreakDown{
				SplitEntityID: info.SplitEntityID,
				Amount:        amount,
			}

		case "RATIO":
			if !tracker.RatioOpenBalance.Set {
				tracker.RatioOpenBalance.Balance = tracker.Balance
				tracker.RatioOpenBalance.Set = true
				tracker.RatioOpenBalance.Total = GetNumRatio(splitInfo)
			}

			ratio := info.SplitValue / tracker.RatioOpenBalance.Total
			amount := ratio * tracker.RatioOpenBalance.Balance
			tracker.Balance -= amount
			response.SplitBreakDown[index] = SplitBreakDown{
				SplitEntityID: info.SplitEntityID,
				Amount:        amount,
			}
		}
		if tracker.Balance <= -1 {
			log.Println(tracker.Balance)
			log.Println("Balance Is Less Than 0")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if response.SplitBreakDown[index].Amount < 0 {
			log.Println(tracker.Balance)
			log.Println(response.SplitBreakDown[index].Amount)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	response.ID = request.ID
	response.Balance = tracker.Balance

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}
