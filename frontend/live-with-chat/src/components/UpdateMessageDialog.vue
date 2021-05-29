<template>
  <div class="text-center">
    <v-btn
      small
      depressed
      @click.stop="show"
    >
      <v-icon small>mdi-pencil-outline</v-icon>
    </v-btn>
    <v-dialog
      v-model="dialog"
      width="500"
      open-on-focus
    >
      <v-card>
        <v-card-title class="headline grey lighten-2">
          New Text
        </v-card-title>
        <v-card-text
          align="center"
        >
          <v-text-field
            v-model="message"
            label="update message"
            clearable
          ></v-text-field>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            text
            color="primary"
            @click="update"
          >
            Update
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  name: 'UpdateMessageDialog',
  props: [
    'id',
    'oldMessage',
  ],
  data () {
    return {
      dialog: false,
      message: "",
    }
  },
  mounted() {
    this.message = this.oldMessage
  },
  methods: {
    show() {
      this.dialog = true
    },
    update() {
      let updateForm = new FormData()
      updateForm.append("id", this.id)
      updateForm.append("text", this.message)

      this.$api.v1.chat.update(
        updateForm
      ).then(() => {
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