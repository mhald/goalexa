package goalexa

//
//
// AudioPlayer interface

const (
	RequestTypeAudioPlayerPlaybackStarted        RequestType = "AudioPlayer.PlaybackStarted"
	RequestTypeAudioPlayerPlaybackFinished       RequestType = "AudioPlayer.PlaybackFinished"
	RequestTypeAudioPlayerPlaybackStopped        RequestType = "AudioPlayer.PlaybackStopped"
	RequestTypeAudioPlayerPlaybackNearlyFinished RequestType = "AudioPlayer.PlaybackNearlyFinished"
	RequestTypeAudioPlayerPlaybackFailed         RequestType = "AudioPlayer.PlaybackFailed"
)

type AudioPlayerContext struct {
	PlayerActivity AudioPlayerActivity `json:"playerActivity"`
}

type AudioPlayerActivity string

const (
	AudioPlayerActivityUnspecified    = ""
	AudioPlayerActivityIdle           = "IDLE"
	AudioPlayerActivityPlaying        = "PLAYING"
	AudioPlayerActivityPaused         = "PAUSED"
	AudioPlayerActivityFinished       = "FINISHED"
	AudioPlayerActivityBufferUnderrun = "BUFFER_UNDERRUN"
)

type AudioPlayerPlayBehavior string

const (
	AudioPlayerPlayBehaviorUnspecified     AudioPlayerPlayBehavior = ""
	AudioPlayerPlayBehaviorReplaceAll      AudioPlayerPlayBehavior = "REPLACE_ALL"
	AudioPlayerPlayBehaviorEnqueue         AudioPlayerPlayBehavior = "ENQUEUE"
	AudioPlayerPlayBehaviorReplaceEnqueued AudioPlayerPlayBehavior = "REPLACE_ENQUEUED"
)

type AudioPlayerClearQueueBehavior string

const (
	AudioPlayerClearQueueBehaviorUnspecified   AudioPlayerClearQueueBehavior = ""
	AudioPlayerClearQueueBehaviorClearEnqueued AudioPlayerClearQueueBehavior = "CLEAR_ENQUEUED"
	AudioPlayerClearQueueBehaviorClearAll      AudioPlayerClearQueueBehavior = "CLEAR_ALL"
)

func CreateDirectiveAudioPlayerPlay(
	behavior AudioPlayerPlayBehavior,
	streamUrl string,
	token string,
	prevToken *string,
	offsetMs int,
) any {
	streamObj := map[string]any{
		"url":                  streamUrl,
		"token":                token,
		"offsetInMilliseconds": offsetMs,
	}
	if prevToken != nil {
		streamObj["expectedPreviousToken"] = *prevToken
	}
	return map[string]any{
		"type":         "AudioPlayer.Play",
		"playBehavior": behavior,
		"audioItem": map[string]any{
			"stream": streamObj,
		},
	}
}

func CreateDirectiveAudioPlayerStop() any {
	return map[string]any{
		"type": "AudioPlayer.Stop",
	}
}

func AudioPlayerClearQueue(
	clearBehavior AudioPlayerClearQueueBehavior,
) any {
	return map[string]any{
		"type":          "AudioPlayer.ClearQueue",
		"clearBehavior": clearBehavior,
	}
}
