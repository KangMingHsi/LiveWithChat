<template>
  <v-dialog
    v-model="showUpload"
    width="500"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
        id="upload-btn"
        class="mx-2"
        fab
        dark
        color="indigo"
        v-bind="attrs"
        v-on="on"
      >
        <v-icon dark>
          mdi-plus
        </v-icon>
      </v-btn>
    </template>

    <div class="text-center">
      <v-row class="container">
        <v-col class="col"> 
          <v-card
            class="upload"
            width="600"
          >
            <v-card-title>
              <h2>Upload Video</h2>
            </v-card-title>

            <validation-observer ref="observer">
              <v-form @submit.prevent="onSubmit">
                <v-card-text>
                  <validation-provider v-slot="{ errors }" name="Video" rules="required">
                    <v-file-input
                      v-model="video"
                      :error-messages="errors"
                      placeholder="Choose a video"
                      label="Video Input"
                      dense
                      filled
                      outlined
                      accept="video/*"
                    ></v-file-input>
                  </validation-provider>
                  <v-text-field
                    v-model="title"
                    label="Title"
                    dense
                    filled
                    outlined
                  ></v-text-field>
                  <v-text-field
                    v-model="description"
                    label="Description"
                    dense
                    filled
                    outlined
                  ></v-text-field>
                </v-card-text>
                <v-card-actions class="text-center">
                  <v-btn class="signin-btn" type="submit" rounded color="blue-grey" dark>
                    Upload
                  </v-btn>
                </v-card-actions>
              </v-form>
            </validation-observer>
          </v-card>
        </v-col>
      </v-row>
    </div>
  </v-dialog>
</template>

<script>
  import { required } from 'vee-validate/dist/rules'
  import { extend, ValidationProvider, setInteractionMode, ValidationObserver } from 'vee-validate'

  setInteractionMode('eager')

  extend('required', {
    ...required,
    message: '{_field_} can not be empty'
  })

  export default {
    name: 'UploadVideoDialog',
    components: {
      ValidationProvider,
      ValidationObserver
    },
    data() {
      return {
        showUpload: false,
        video: null,
        title: '',
        description: '',
      }
    },
    watch: {
      showUpload: function (val) {
        val || this.clear()
      }
    },
    methods: {
      async onSubmit() {
        const valid = await this.$refs.observer.validate()
        if (valid) {
          let title = this.title != '' ? this.title: this.video.name
          let description = this.description != '' ? this.description: 'created at ' + Date().toLocaleString()
          var video_type = this.video.type
          video_type = "." + video_type.substring(video_type.lastIndexOf("/") + 1)

          var uploadForm = new FormData()
          uploadForm.append("video", this.video)
          uploadForm.append("title", title)
          uploadForm.append("video_type", video_type)
          uploadForm.append("description", description)

          this.$api.v1.stream.upload(
            uploadForm
          ).then((resp) => {
            console.log(resp.status)
            this.clear()
          }).catch((error) => {
            console.log(error)
          })
        }
      },
      clear() {
        this.video = null
        this.title = ''
        this.description = ''
        this.$refs.observer.reset()
      },
    },
  }
</script>

<style lang="scss" scoped>
  .container {
    padding: 0;
    background: rgb(255, 255, 255);
    width: 100%;
    box-shadow: 0 0 1px 1px rgba($color: #000000, $alpha: 0.1);
    box-sizing: border-box;
    .col {
      padding: 0;
      max-width: 1000px;
      box-shadow: 0 0 1px 1px rgba($color: #000000, $alpha: 0.1);
      .upload {
        padding: 0;
        box-sizing: border-box;
        background: #ffffff;
        color: rgb(5, 5, 5);
        h2 {
          text-align: center;
          margin: 0px 0;
        }
        .v-btn {
          width: 100%;
          color: #ffffff;
        }
      }
    }
  }
  #upload-btn {
    position: fixed;
    bottom: 5px;
    right: 0;
  }
</style>