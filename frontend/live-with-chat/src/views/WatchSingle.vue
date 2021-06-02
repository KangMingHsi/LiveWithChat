<template>
  <v-container>
    <div class="video-section">
      <video
        class="video-player"
        controls
        type="application/x-mpegURL"
        :src="src"
      ></video>
    </div>
    <v-row class="non-video">
      <v-col>
        <p>{{ title }}</p>
        <p>{{ description }}</p>
        <v-divider></v-divider>
      </v-col>

      <v-col>
        <v-card
          elevation="16"
          max-width="400"
          class="mx-auto"
          max-height="500"
        >
          <v-virtual-scroll
            :items="messages"
            :item-height="50"
            height="400"
          >
            <template v-slot:default="{ item }">
              <v-list-item
                @mouseenter="hoverOn(item)"
                @mouseleave="hoverOff(item)"
              >
                <v-list-item-icon>
                  <v-icon>mdi-account-circle</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>
                    {{ item.Name }} {{ item.Text }}
                  </v-list-item-title>
                </v-list-item-content>
                <v-list-item-action v-show="isOwnMessage(item)">
                  <UpdateMessageDialog
                    :id="item.ID"
                    :oldMessage="item.Text"
                    @reload="getMessages"
                  />
                </v-list-item-action>
                <v-list-item-action v-show="isOwnMessage(item)">
                  <DeleteMessageDialog
                    :id="item.ID"
                    @reload="getMessages"
                  />
                </v-list-item-action>
              </v-list-item>
            </template>
          </v-virtual-scroll>
          
          <v-divider></v-divider>
          <v-row
            dense
            align="center"
          >
            <v-col
              cols="10"
            >
              <v-text-field
                v-model="message"
                label="send message"
                clearable
                dense
                style="padding-left: 5%;"
              ></v-text-field>
            </v-col>
            <v-col
              cols="2"
            >
              <v-btn
                icon
                color="blue"
                @click="sendMessage"
              >
                <v-icon>mdi-send</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import UpdateMessageDialog from '@/components/UpdateMessageDialog';
import DeleteMessageDialog from '@/components/DeleteMessageDialog'
export default {
  name: 'WatchSingle',
  components: {
    UpdateMessageDialog,
    DeleteMessageDialog
  },
  data() {
    return {
      message: "",
      watch: {},
      messages: [],
      id: "",
      title: "",
      description: "",
      src: "",
    }
  },
  mounted() {
    this.id = this.$route.params.id
    this.title = this.$store.state.watch.title
    this.description = this.$store.state.watch.description
    this.messages = []
    this.src = "/static/" + this.id + "/playlist.m3u8"

    this.getMessages()
  },
  methods: {
    sendMessage() {
      var sendForm = new FormData()
      sendForm.append("text", this.message)
      sendForm.append("vid", this.id)

      this.$api.v1.chat.send(
        sendForm
      ).then(() => {
        this.message = ""
      })
    },
    isOwnMessage(item) {
      return item.IsHover && item.OwnerID == this.$store.state.auth.id
    },
    hoverOn(item) {
      item.IsHover = true
    },
    hoverOff(item) {
      item.IsHover = false
    },
    getMessages() {
      this.messages = []
      this.$api.v1.chat.list({
        vid: this.id
      }).then((resp) => {
        let data = resp.data
        data.forEach(element => {
          element.Name = 'Sam'
          element.IsHover = false
          this.messages.push(element)
        });
      }).catch((err) => {
        console.log(err)
      })
    }
  },
}
</script>
<style lang="scss" scoped>
  .video-section {
    width: 100%;
    height: 480px;
  }
  .non-video {
    padding-top: 20%;
  }
  .watch-single {
    margin-top: 30px;
  }
  .hero {
    margin-bottom: 70px;
  }
  .watch-images {
    margin-top: 50px;
  }
  .description {
    margin-bottom: 30px;
  }
</style>