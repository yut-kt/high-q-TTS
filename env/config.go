package env

import "os"

var (
	RootDir                  = os.Getenv("GOPATH") + "/src/github.com/yut-kt/high-q-TTS"
	OrigDataDir              = RootDir + "/data/dourokoutsuu"
	AllBodyPath              = RootDir + "/data/all_body.txt"
	AllIndexedBodyPath       = RootDir + "/data/all_indexed_body.txt"
	AllUniqueBodyPath        = RootDir + "/data/all_unique_body.txt"
	AllIndexedUniqueBodyPath = RootDir + "/data/all_indexed_unique_body.txt"
	ClusterBaseDir           = RootDir + "/data/cluster/base"
	ClusterClassifiedDir     = RootDir + "/data/cluster/classified"
	
	DataBaseDriver = "sqlite3"
	DataBase       = RootDir + "/db/high_q_TTS.sqlite"
)

func init() {
	if os.Getenv("TEST") == "true" {
		OrigDataDir = RootDir + "/data/test_data"
	}
}
