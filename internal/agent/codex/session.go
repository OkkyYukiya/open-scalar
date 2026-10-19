package codex

import "time"

// Session は Codex の 1 セッション(rollout-*.jsonl 1 ファイル)の計測結果を表す。
type Session struct {
	Path          string    // jsonl ファイルの絶対パス
	ID            string    // session_meta の session_id
	ThreadName    string    // session_index.jsonl 由来。無い場合は空
	Cwd           string    // session_meta の cwd
	SizeBytes     int64     // ファイルサイズ(os.FileInfo.Size() が int64)
	Lines         int       // 行数
	Turns         int       // task_started の回数
	TotalTokens   int64     // 最後の token_count の total_tokens
	ContextWindow int64     // model_context_window
	CompactCount  int       // compacted 行の数
	UpdatedAt     time.Time // ファイルの最終更新時刻
}

func (s Session)  ContextRatio () float64 {
	if s.ContextWindow == 0 {
		return 0
	}
	return float64(s.TotalTokens) / float64(s.ContextWindow)
}