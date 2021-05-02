<template>
  <v-dialog
    v-model="showUpdate"
    width="500"
  >
    <v-card
      class="update-card"
      width="600"
    >
      <v-card-title>Update video</v-card-title>
      <v-form @submit.prevent="executeUpdate">
        <v-card-text>
          <v-text-field
            v-model="handledWatch.ID"
            label="ID"
            dense
            filled
            outlined
            readonly
          ></v-text-field>
          <v-text-field
            v-model="handledWatch.Title"
            label="Title"
            dense
            filled
            outlined
          ></v-text-field>
          <v-text-field
            v-model="handledWatch.Description"
            label="Description"
            dense
            filled
            outlined
          ></v-text-field>
        </v-card-text>
        <v-card-actions class="text-center">
          <v-spacer></v-spacer>
          <v-btn
            class="update-btn"
            type="submit"
            text
          >
            Update
          </v-btn>
          <v-btn
            text
            @click="showUpdate = false"
          >
            Cancel
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'UpdateVideoDialog',
  data() {
    return {
      handledWatch: {},
      showUpdate: false,
    }
  },
  methods: {
    show(handledWatch) {
      this.handledWatch = handledWatch
      this.showUpdate = true
    },
    executeUpdate() {
      this.showUpdate = false

      var updateForm = new FormData()
      updateForm.append("vid", this.handledWatch.ID)
      updateForm.append("title", this.handledWatch.Title)
      updateForm.append("description", this.handledWatch.Description)
      this.$api.v1.stream.update(
        updateForm
      ).then(() => {
        this.$emit('get-videos')
      })
    },
  },
}
</script>

<style>

</style>