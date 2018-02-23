package env

import "os"

var (
	RootDir = os.Getenv("GOPATH") + "/src/github.com/yut-kt/high-q-TTS"

	dataDir                   = RootDir + "/data"
	OrigDataDir               = dataDir + "/dourokoutsuu"
	AllBodyPath               = dataDir + "/all_body.txt"
	AllIndexedBodyPath        = dataDir + "/all_indexed_body.txt"
	AllUniqueBodyPath         = dataDir + "/all_unique_body.txt"
	AllIndexedUniqueBodyPath  = dataDir + "/all_indexed_unique_body.txt"
	AllPhraseDivisionBodyPath = dataDir + "/all_phrase_division_body.txt"

	clusterBaseDir           = RootDir + "/data/cluster/base"
	ClusterNormalDir         = clusterBaseDir + "/normal"
	ClusterPhraseDivisionDir = clusterBaseDir + "/phrase_division"
	ClusterClassifiedDir     = RootDir + "/data/cluster/classified"

	DataBaseDriver = "sqlite3"
	DataBase       = RootDir + "/db/high_q_TTS.sqlite"
)

func init() {
	if os.Getenv("TEST") == "true" {
		OrigDataDir = RootDir + "/data/test_data"
	}
}
