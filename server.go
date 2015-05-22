package goblog

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

//Server ...
func Server(port int) {
	log.Info("Server start. Please visit http://localhost:" + strconv.Itoa(port))
	log.Infoln("Press ctrl-c to stop")
	if err := http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(config.PublicDir))); err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
	}
	return
}
