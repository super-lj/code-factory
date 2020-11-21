package mock

// define mock data stuctures
type RepoInfo struct {
	Name        string
	BranchNames []string
	CommitHashs []string
	MaxRunNum   int32
}

type BranchInfo struct {
	Name       string
	CommitHash string
	RunNums    []int32
}

type CommitInfo struct {
	Hash    string
	Msg     string
	Author  string
	RunNums []int32
}

type RunInfo struct {
	Num            int32
	StartTimestamp int32
	Duration       int32
	Status         string
	Log            string
	BranchName     string
	CommitHash     string
}

// mock DB tables
var mockRepoDB = []RepoInfo{
	{
		Name:        "test_repo_a",
		BranchNames: []string{"main", "dev", "dev_2"},
		CommitHashs: []string{"01df63c", "29eea5d", "f26fbd4"},
		MaxRunNum:   3,
	},
	{
		Name:        "test_repo_b",
		BranchNames: []string{"main", "dev_b", "dev_b_2"},
		CommitHashs: []string{"1741c90", "1ff945d", "55347bc"},
		MaxRunNum:   3,
	},
}

var mockBranchDB = map[string][]BranchInfo{
	"test_repo_a": {
		{
			Name:       "main",
			CommitHash: "01df63c",
			RunNums:    []int32{3},
		},
		{
			Name:       "dev",
			CommitHash: "29eea5d",
			RunNums:    []int32{2},
		},
		{
			Name:       "dev_2",
			CommitHash: "f26fbd4",
			RunNums:    []int32{1},
		},
	},
	"test_repo_b": {
		{
			Name:       "main",
			CommitHash: "1741c90",
			RunNums:    []int32{3},
		},
		{
			Name:       "dev_b",
			CommitHash: "1ff945d",
			RunNums:    []int32{2},
		},
		{
			Name:       "dev_b_2",
			CommitHash: "55347bc",
			RunNums:    []int32{1},
		},
	},
}

var mockCommitDB = map[string][]CommitInfo{
	"test_repo_a": {
		{
			Hash:    "01df63c",
			Msg:     "Commit 01df63c Msg.",
			Author:  "Peixuan Li",
			RunNums: []int32{3},
		},
		{
			Hash:    "29eea5d",
			Msg:     "Commit 29eea5d Msg.",
			Author:  "Peixuan Li",
			RunNums: []int32{2},
		},
		{
			Hash:    "f26fbd4",
			Msg:     "Commit f26fbd4 Msg.",
			Author:  "Peixuan Li",
			RunNums: []int32{1},
		},
	},
	"test_repo_b": {
		{
			Hash:    "1741c90",
			Msg:     "Commit 1741c90 Msg.",
			Author:  "Xingyou Ji",
			RunNums: []int32{3},
		},
		{
			Hash:    "1ff945d",
			Msg:     "Commit 1ff945d Msg.",
			Author:  "Xingyou Ji",
			RunNums: []int32{2},
		},
		{
			Hash:    "55347bc",
			Msg:     "Commit 55347bc Msg.",
			Author:  "Xingyou Ji",
			RunNums: []int32{1},
		},
	},
}

var mockRunDB = map[string][]RunInfo{
	"test_repo_a": {
		{
			Num:            3,
			StartTimestamp: 1605053500,
			Duration:       142,
			Status:         "IN_PROGRESS",
			BranchName:     "main",
			CommitHash:     "01df63c",
			Log: `go run xxxxxx
branch main commit 01df63c run 3 succeed!`,
		},
		{
			Num:            2,
			StartTimestamp: 1605053200,
			Duration:       212,
			Status:         "SUCCEED",
			BranchName:     "dev",
			CommitHash:     "29eea5d",
			Log:            ``,
		},
		{
			Num:            1,
			StartTimestamp: 1605053000,
			Duration:       123,
			Status:         "FAILED",
			BranchName:     "dev_2",
			CommitHash:     "f26fbd4",
			Log:            ``,
		},
	},
	"test_repo_b": {
		{
			Num:            3,
			StartTimestamp: 1605053500,
			Duration:       142,
			Status:         "IN_PROGRESS",
			BranchName:     "main",
			CommitHash:     "1741c90",
			Log: `go run xxxxxx
branch main commit 1741c90 run 3 succeed!`,
		},
		{
			Num:            2,
			StartTimestamp: 1605053200,
			Duration:       212,
			Status:         "SUCCEED",
			Log:            ``,
			BranchName:     "dev_b",
			CommitHash:     "1ff945d",
		},
		{
			Num:            1,
			StartTimestamp: 1605053000,
			Duration:       123,
			Status:         "FAILED",
			Log:            ``,
			BranchName:     "dev_b_2",
			CommitHash:     "55347bc",
		},
	},
}

// mock Cache by map
var mockRepoCache = map[string]RepoInfo{}
var mockBranchCache = map[string]map[string]BranchInfo{}
var mockCommitCache = map[string]map[string]CommitInfo{}
var mockRunCache = map[string]map[int32]RunInfo{}

func init() {
	// init Repo Cache
	for _, repo := range mockRepoDB {
		mockRepoCache[repo.Name] = repo
	}
	// init Branch Cache
	for repoName, branches := range mockBranchDB {
		mockBranchCache[repoName] = make(map[string]BranchInfo)
		for _, br := range branches {
			mockBranchCache[repoName][br.Name] = br
		}
	}
	// init Commit Cache
	for repoName, commits := range mockCommitDB {
		mockCommitCache[repoName] = make(map[string]CommitInfo)
		for _, c := range commits {
			mockCommitCache[repoName][c.Hash] = c
		}
	}
	// init Run Cache
	for repoName, runs := range mockRunDB {
		mockRunCache[repoName] = make(map[int32]RunInfo)
		for _, r := range runs {
			mockRunCache[repoName][r.Num] = r
		}
	}
}

// mock RPC interfaces
func GetRepoNames() []string {
	var res = []string{}
	for _, repo := range mockRepoDB {
		res = append(res, repo.Name)
	}
	return res
}

func GetRepoInfo(name string) *RepoInfo {
	repo, ok := mockRepoCache[name]
	if !ok {
		return nil
	}
	return &repo
}

func GetBranchInfo(repoName string, branchName string) *BranchInfo {
	branchsMap, repoNameOK := mockBranchCache[repoName]
	if !repoNameOK {
		return nil
	}
	branch, branchNameOK := branchsMap[branchName]
	if !branchNameOK {
		return nil
	}
	return &branch
}

func GetCommitInfo(repoName string, commitHash string) *CommitInfo {
	commitsMap, repoNameOK := mockCommitCache[repoName]
	if !repoNameOK {
		return nil
	}
	commit, commmitHashOK := commitsMap[commitHash]
	if !commmitHashOK {
		return nil
	}
	return &commit
}

func GetRunInfo(repoName string, runNum int32) *RunInfo {
	runsMap, repoNameOK := mockRunCache[repoName]
	if !repoNameOK {
		return nil
	}
	run, runNumOK := runsMap[runNum]
	if !runNumOK {
		return nil
	}
	return &run
}
