module cryptotracker

go 1.17

replace krakenApi => ../krakenApi

require (
	krakenApi v0.0.0-00010101000000-000000000000
	pairReader v0.0.0-00010101000000-000000000000
)

replace pairReader => ../pairReader
