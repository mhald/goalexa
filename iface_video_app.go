package goalexa

//
//
// VideoApp interface

func CreateDirectiveVideoAppLaunch(
	streamUrl string,
	title string,
	subtitle string,
) map[string]any {
	videoItemObj := map[string]any{
		"source": streamUrl,
	}
	if title != "" || subtitle != "" {
		videoItemObj["metadata"] = map[string]any{
			"title":    title,
			"subtitle": subtitle,
		}
	}
	return map[string]any{
		"type":      "VideoApp.Launch",
		"videoItem": videoItemObj,
	}
}
