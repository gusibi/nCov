package main

import (
	"github.com/gusibi/nCov/internal/data_sync"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func main() {
	cloudfunction.Start(data_sync.Run)
}
