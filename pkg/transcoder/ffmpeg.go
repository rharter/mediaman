package transcoder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	//FFMPEG_ARGS_HLS string = " -async 1 -ss 00:00:05 -acodec libfdk_aac -vbr 3 -b:v \\3000k -ac 2 -vcodec libx264 -preset superfast -tune zerolatency  -threads 2 -s 1280x720 -flags -global_header -map 0:0 -map 0:1 -f segment -segment_time 10 -segment_list index.m3u8 -segment_format mpegts -segment_wrap 10 -segment_list_size 6 -segment_list_flags live hls/%03d.ts"
	FFMPEG_ARGS_HLS     string = " -bsf h264_mp4toannexb -f segment -codec copy -map 0 -segment_time 2 -segment_format mpegts -segment_list_flags +live -segment_list_type m3u8 -individual_header_trailer 1 -segment_list index.m3u8 %09d.ts"
	FFMPEG_ARGS_DASH    string = " "
	FFMPEG_ARGS_MP4     string = " -codec copy video.mp4"
	FFMPEG_ARG_WEBM     string = " -f webm -codec:v libvpx -quality realtime -cpu-used 0 -b:v 1200k -qmin 10 -qmax 42 -minrate 1200k -maxrate 1200k -bufsize 1500k -threads 1 -codec:a libvorbis -b:a 128k video.webm "
	FFMPEG_ARG_WEBM_HLS string = " -f webm -force_key_frames expr:gte(t,n_forced*2) -codec:v libvpx -quality realtime -cpu-used 0 -b:v 1200k -qmin 10 -qmax 42 -maxrate 1200k -bufsize 1500k -lag-in-frames 0 -rc_lookahead 0 -flags +global_header -codec:a libvorbis -b:a 128k -flags +global_header -map 0 -f segment -segment_list_flags +live -segment_time 2 -segment_format webm -flags +global_header -segment_list webm_index.m3u8 webm/%09d.webm "
)

const (
	FFMPEG_CMD_START_PROD string = "-y -v quiet "
	FFMPEG_CMD_START_DEV  string = "-y -v debug "
	FFMPEG_CMD_INPUT      string = " -i pipe:0 "
)

func (t *TranscodeSession) createOutputDirectories() error {

	// HLS: <hls-data-path>/<id>/hls
	p := filepath.Join(t.OutputDir, t.ID, "hls")
	err := os.MkdirAll(p, 0775)
	if err != nil {
		log.Printf("Failed to create directory %s: %s", p, err)
		return err
	}

	return err
}

// assemble transcode command
func (t *TranscodeSession) buildTranscodeCommand() string {
	var cmd string

	cmd = FFMPEG_CMD_START_DEV

	cmd += fmt.Sprintf(" -i %s ", t.InputFile)

	// Just HLS for now
	cmd += FFMPEG_ARG_WEBM_HLS

	return cmd
}
