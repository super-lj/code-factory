namespace go ci

struct RepoInfo {
    1: string name,
    2: list<string> branchNames,
    3: list<string> commitHashs,
    4: i32 maxRunNum,
}

struct BranchInfo {
    1: string name,
    2: string commitHash,
    3: list<i32> runNums,
}

struct CommitInfo {
    1: string hash,
    2: string msg,
    3: string author,
    4: list<i32> runNums,
}

struct RunInfo {
    1: i32 num,
    2: i32 startTimestamp,
    3: i32 duration,
    4: string status,
    5: string log,
    6: string branchName,
    7: string commitHash,
}

service CIBackend {
    list<string> getRepoNames(),

    RepoInfo getRepoInfo(1:string name),

    BranchInfo getBranchInfo(1:string repoName, 2:string branchName),

    CommitInfo getCommitInfo(1:string repoName, 2:string commitHash),

    RunInfo getRunInfo(1:string repoName, 2:i32 runNum),
}
