package oxygengo

type Oxygen struct {
  CName string
  RName string
  GName string
  NWname string
  CUsage float64
  RUsage float64
  GUsage float64
  NWUsage float64
  CSys float64
  RSys float64
  GSys float64
  NWSys float64
  CUsr float64
  RUsr float64
  GUsr float64
  NWUsr float64
  OTask OxygenTask
  Tasks []OxygenTask{}
}

type OxygenTask struct {
  Name string
  PID int
  CImpact float64
  RImpact float64
  GImpact float64
  NWImpact float64
}

func init() {
  
}

func GetOxygen() {
  return Oxygen{}
}
