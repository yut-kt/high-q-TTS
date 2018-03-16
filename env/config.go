package env

import "os"

var (
	RootDir = os.Getenv("GOPATH") + "/src/github.com/yut-kt/high-q-TTS"

	dataDir = RootDir + "/data"

	OrigDataDir = dataDir + "/dourokoutsuu"

	allDir                    = dataDir + "/all"
	AllBodyPath               = allDir + "/body.txt"
	AllIndexedBodyPath        = allDir + "/indexed_body.txt"
	AllUniqueBodyPath         = allDir + "/unique_body.txt"
	AllIndexedUniqueBodyPath  = allDir + "/indexed_unique_body.txt"
	AllPhraseDivisionBodyPath = allDir + "/phrase_division_body.txt"

	clusterDir               = RootDir + "/data/cluster"
	ClusterNormalDir         = clusterDir + "/base/normal"
	ClusterPhraseDivisionDir = clusterDir + "/base/phrase_division"
	ClusterClassifiedDir     = clusterDir + "/classified"

	DataBaseDriver = "sqlite3"
	DataBase       = RootDir + "/db/high_q_TTS.sqlite"
)

func init() {
	if os.Getenv("TEST") == "true" {
		OrigDataDir = RootDir + "/data/test_data"
	}
}
