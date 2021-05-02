<template>
  <v-dialog
    v-model="showDeletion"
    width="500"
  >
    <v-card
      class="delete-card"
      width="600"
      :key="handledWatch.ID"
    >
      <v-card-title>Delete video</v-card-title>
      <v-card-text>It's going to delete video #{{ handledWatch.ID }}.</v-card-text>
      <v-card-actions class="text-center">
        <v-spacer></v-spacer>
        <v-btn
          text
          @click="executeDeletion"
        >
          Yes
        </v-btn>
        <v-btn
          text
          @click="showDeletion = false"
        >
          No
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'DeleteVideoDialog',
  data() {
    return {
      handledWatch: {},
      showDeletion: false,
    }
  },
  methods: {
    show(handledWatch) {
      this.handledWatch = handledWatch
      this.showDeletion = true
    },
    executeDeletion() {
      this.showDeletion = false

      let query = {'vid': this.handledWatch.ID}
      this.$api.v1.stream.delete(
        query
      ).then(() => {
        this.$emit('get-videos')
      })
    }
  },
}
</script>

<style>

</style>