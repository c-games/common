package fail

import (
	"log"
	"runtime"
)


func getFrame(skipFrames int) runtime.Frame {
    // We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
    targetFrameIndex := skipFrames + 2

    // Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
    programCounters := make([]uintptr, targetFrameIndex+2)
    n := runtime.Callers(0, programCounters)

    frame := runtime.Frame{Function: "unknown"}
    if n > 0 {
        frames := runtime.CallersFrames(programCounters[:n])
        for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
            var frameCandidate runtime.Frame
            frameCandidate, more = frames.Next()
            if frameIndex == targetFrameIndex {
                frame = frameCandidate
            }
        }
    }

    return frame
}

func FailOnError(err error, msg string) {
	if err != nil {
		frame := getFrame(2)
		fn := frame.Function
		file := frame.File
		line := frame.Line
		log.Fatalf("[file:%s, %v, fn: %s]  %s: %s",
			file, line, fn,
			msg, err)
	}
}
