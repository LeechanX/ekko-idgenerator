package generator

/**
|------------------------------------------------------------------------------|
|                     Ekko ID generator design                                 |
|------------------------------------------------------------------------------|
| 64bits                                                                       |
| : 39bits for timestamp                                                       |
| : 3bits for idc                                                              |
| : 8~12bits for workerID                                                      |
| : 0~4bits for products                                                       |
| : 10bits for currency                                                        |
|------------------------------------------------------------------------------|
|                                                                              |
| @if 8bits for workerID                                                       |
|  [.......................................][...][........][....][..........]  |
|  [                timestamp              ][idc][workerID][prod][ currency ]  |
|                                                                              |
| @if 12bits for workerID, no product                                          |
|  [.......................................][...][........][....][..........]  |
|  [                timestamp              ][idc][   workerID   ][ currency ]  |
|                                                                              |
|------------------------------------------------------------------------------|
*/

const (
	idcOffset                 = 22
	idcUpperBound             = 8
	maxMultiGetCount          = 1024
	productUpperBound         = 16
	productOffset             = 10
	workerIDOffsetBase        = 10
	TimestampOffset           = 25
	CurrencyBound             = 1024
	CurrencyMask       uint64 = 0x3FF
	mmapFilePathTpl           = "/opt/ekko_file_%d"
)
