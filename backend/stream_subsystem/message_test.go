package stream_subsystem

import "testing"

func TestNewMessage(t *testing.T) {
	var (
		factory = NewMessageFactory(0)
		text = "HI"
		videoID = "video"
		ownerID = "owner"
	)

	message := factory.NewMessage(text, videoID, ownerID)
	if message.ID != 0 || message.Text != text ||
		message.VideoID != videoID || message.OwnerID != ownerID{

		t.Errorf("ID should be 0")
	}
}