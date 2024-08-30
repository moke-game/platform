package room

import roompb "github.com/moke-game/platform/api/gen/room/api"

type Frames struct {
	frameIndex uint32
	frames     map[uint32]*roompb.FrameData
}

func NewFrames() *Frames {
	return &Frames{
		frames: make(map[uint32]*roompb.FrameData),
	}
}

func (f *Frames) tick() {
	f.frameIndex++
}

func (f *Frames) getFrameIndex() uint32 {
	return f.frameIndex
}

func (f *Frames) checkCmdIsExist(cmd *roompb.CmdData) bool {
	if frame, ok := f.frames[f.frameIndex]; ok {
		for _, c := range frame.Cmds {
			if c.GetUid() == cmd.GetUid() {
				return true
			}
		}
	}
	return false
}

func (f *Frames) PushCmd(cmd *roompb.CmdData) bool {
	if frame, ok := f.frames[f.frameIndex]; ok {
		if f.checkCmdIsExist(cmd) {
			return false
		}
		frame.Cmds = append(frame.Cmds, cmd)
	} else {
		f.frames[f.frameIndex] = &roompb.FrameData{
			FrameIndex: f.frameIndex,
			Cmds:       []*roompb.CmdData{cmd},
		}
	}
	return true
}

func (f *Frames) getRangeFrames(start, end uint32) []*roompb.FrameData {
	var frames []*roompb.FrameData
	for i := start; i <= end; i++ {
		if frame, ok := f.frames[i]; ok {
			frames = append(frames, frame)
		}
	}
	return frames
}

func (f *Frames) getCurrentFrame() *roompb.FrameData {
	return f.frames[f.frameIndex]
}
