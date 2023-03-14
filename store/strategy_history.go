package store

import (
	"encoding/json"
	"github.com/eolinker/apinto-dashboard/entry/history-entry"
	"github.com/eolinker/apinto-dashboard/entry/strategy-entry"
)

type IStrategyHistoryStore interface {
	BaseHistoryStore[strategy_entry.StrategyHistory]
}

type strategyHistoryHandler struct {
}

func (s *strategyHistoryHandler) Kind() string {
	return "strategy"
}

func (s *strategyHistoryHandler) Decode(r *history_entry.History) *strategy_entry.StrategyHistory {
	oldValue := new(strategy_entry.StrategyHistoryInfo)
	_ = json.Unmarshal([]byte(r.OldValue), oldValue)
	newValue := new(strategy_entry.StrategyHistoryInfo)
	_ = json.Unmarshal([]byte(r.NewValue), newValue)
	history := &strategy_entry.StrategyHistory{
		Id:          r.Id,
		NamespaceId: r.NamespaceID,
		StrategyId:  r.TargetID,
		OptTime:     r.OptTime,
		OptType:     r.OptType,
		OldValue:    *oldValue,
		NewValue:    *newValue,
		Operator:    r.Operator,
	}
	return history
}

func newStrategyHistoryStore(db IDB) IStrategyHistoryStore {
	var historyHandler DecodeHistory[strategy_entry.StrategyHistory] = new(strategyHistoryHandler)
	return CreateHistory(historyHandler, db, history_entry.HistoryKindStrategy)
}
