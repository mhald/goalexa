package goalexa

//
//
// Interface: AudioPlayer

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

type AudioItemMetadata struct {
	Title           string        `json:"title,omitempty"`
	Subtitle        string        `json:"subtitle,omitempty"`
	Art             *DisplayImage `json:"art,omitempty"`
	BackgroundImage *DisplayImage `json:"backgroundImage,omitempty"`
}

type AudioItemCaptionDataType string

const (
	AudioItemCaptionDataTypeUnspecified AudioItemCaptionDataType = ""
	AudioItemCaptionDataTypeWebvtt      AudioItemCaptionDataType = "WEBVTT"
)

type AudioItemCaptionData struct {
	Type    AudioItemCaptionDataType `json:"type,omitempty"`
	Content string                   `json:"content,omitempty"`
}

type AudioItemStream struct {
	Url                   string                `json:"url"`
	Token                 string                `json:"token"`
	ExpectedPreviousToken string                `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  uint64                `json:"offsetInMilliseconds"`
	CaptionData           *AudioItemCaptionData `json:"captionData,omitempty"`
}

type AudioPlayerAudioItem struct {
	Stream   AudioItemStream    `json:"stream"`
	Metadata *AudioItemMetadata `json:"metadata,omitempty"`
}

//
//
//

const (
	DirectiveTypeAudioPlayerPlay       DirectiveType = "AudioPlayer.Play"
	DirectiveTypeAudioPlayerStop       DirectiveType = "AudioPlayer.Stop"
	DirectiveTypeAudioPlayerClearQueue DirectiveType = "AudioPlayer.ClearQueue"
)

func CreateDirectiveAudioPlayerPlay(
	behavior AudioPlayerPlayBehavior,
	streamUrl string,
	token string,
	prevToken *string,
	offsetMs uint64,
) *Directive {
	streamObj := AudioItemStream{
		Url:                  streamUrl,
		Token:                token,
		OffsetInMilliseconds: offsetMs,
	}
	if prevToken != nil {
		streamObj.ExpectedPreviousToken = *prevToken
	}
	return &Directive{
		Type:         DirectiveTypeAudioPlayerPlay,
		PlayBehavior: behavior,
		AudioItem: &AudioPlayerAudioItem{
			Stream: streamObj,
		},
	}
}

func CreateDirectiveAudioPlayerStop() *Directive {
	return &Directive{
		Type: DirectiveTypeAudioPlayerStop,
	}
}

func AudioPlayerClearQueue(
	clearBehavior AudioPlayerClearQueueBehavior,
) *Directive {
	return &Directive{
		Type:          DirectiveTypeAudioPlayerClearQueue,
		ClearBehavior: clearBehavior,
	}
}
