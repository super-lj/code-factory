namespace go ci

struct TrafficEnv {
    1: bool Open = false,
    2: string Env = "",
}

struct Base {
    1: string LogID = "",
    2: string Caller = "",
    3: string Addr = "",
    4: string Client = "",
    5: optional TrafficEnv TrafficEnv,
    6: optional map<string, string> Extra,
}

struct BaseResp {
    1: string StatusMessage = "",
    2: i32 StatusCode = 0,
    3: optional map<string, string> Extra,
}

struct IsTargetRepoUpdatedRequest {
    1:      optional i64 RepoId,
    255:    Base Base
}

struct IsTargetRepoUpdatedResponse {
    1:      optional bool IsUpdated,
    255:    BaseResp BaseResp
}

struct FetchTargetRepoLastCommitRequest {
    1:      optional i64 RepoId,
    255:    Base Base
}

struct CommitStruct {
    1: optional i64 Id,
    2: optional string  Msg,
    3: optional string Author,
    4: optional i64 LastUpdateTime
}

struct FetchTargetRepoLastCommitResonse {
    1:      optional CommitStruct Commit,
    255:    BaseResp BaseResp
}

service CIService {
    IsTargetRepoUpdatedResponse         IsTargetRepoUpdated         (1: IsTargetRepoUpdatedRequest req),
    FetchTargetRepoLastCommitResonse    FetchTargetRepoLastCommit   (1: FetchTargetRepoLastCommitRequest req)
}
