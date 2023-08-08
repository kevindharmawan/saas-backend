package model

type Subscription struct {
	ID            string
	CustomerID    string
	PlanID        string
	CurrentPeriod int64
	NextPeriod    int64
}
