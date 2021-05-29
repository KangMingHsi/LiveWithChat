<template>
  <div class="text-center">
    <v-btn
      small
      depressed
      @click.stop="show"
    >
      <v-icon small>mdi-trash-can-outline</v-icon>
    </v-btn>

    <v-dialog
      v-model="dialog"
      width="500"
      open-on-focus
    >
      <v-card>
        <v-card-title class="headline grey lighten-2">
          Delete Message
        </v-card-title>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            text
            color="primary"
            @click="dialog = false"
          >
            No
          </v-btn>
          <v-btn
            text
            color="primary"
            @click="remove"
          >
            Yes
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  name: 'DeleteMessageDialog',
  props: [
    'id'
  ],
  data() {
    return {
      dialog: false
    }
  },
  methods: {
    show() {
      this.dialog = true
    },
    remove() {
      this.$api.v1.chat.delete({
        id: this.id
      }).then(() => {
        this.$emit('reload')
      }).finally(() => {
        this.dialog = false
      })
    }
  },
}
</script>

<style>

</style>