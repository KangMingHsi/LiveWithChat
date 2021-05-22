<template>
  <div class="columns is-multiline">
    <div v-for="watch in watches" :key="watch.ID" class="column is-one-quarter">
      <div
        class="watch-card"
      >
        <v-menu>
          <template v-slot:activator="{ on, attrs }">
            <div
              v-bind="attrs"
              @contextmenu.stop="showAction($event, on)"
              @click="to(watch)"
            >
              <WatchCard :watch="watch" />
            </div>
          </template>
          <v-list>
            <v-list-item
              v-for="item in videoItems"
              :key="item.title"
              link
              @click.stop="item.action(watch)"
            >
              <v-list-item-title>
                {{ item.title }}
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </div>
    <div
      class="update-dialog"
    >
      <UpdateVideoDialog
        ref="updateDialog"
        @get-videos="getVideos"
      />
    </div>
    <div
      class="delete-dialog"
    >
      <DeleteVideoDialog
        ref="deleteDialog"
        @get-videos="getVideos"
      />
    </div>
  </div>
</template>
<script>
  import WatchCard from '@/components/WatchCard';
  import UpdateVideoDialog from '@/components/UpdateVideoDialog'
  import DeleteVideoDialog from '@/components/DeleteVideoDialog'
  export default {
    name: 'WatchesList',
    props: [
      'conditions',
      'editable',
    ],
    components : {
      WatchCard,
      UpdateVideoDialog,
      DeleteVideoDialog,
    },
    data () {
      return {
        watches: [],
        showUpdate: false,
        showDeletion: false,
        videoItems: [
          {
            title: 'Update',
            action: (w) => {
              this.$refs.updateDialog.show(w)
            },
          },
          {
            title: 'Delete',
            action: (w) => {
              this.$refs.deleteDialog.show(w)
            },
          },
        ],
      }
    },
    created() {
      this.getVideos()
    },
    methods: {
      getVideos() {
        let query = {}
        if (this.conditions !== undefined && typeof(this.conditions) === Object) {
          for (const key in this.conditions) {
            query[key] = this.conditions[key]
          }
        }
        this.$api.v1.stream.list(
          query
        ).then((resp) => {
          let data = resp.data
          this.watches = data
        })
      },
      to (watch) {
        this.$store.dispatch('watch/setWatch', {
          "description": watch.Description,
          "title": watch.Title,
          "id": watch.ID,
        })

        this.$router.push('/watch/' + watch.ID)
      },
      showAction (event, on) {
        if (this.editable) {
          on.click(event)
          event.preventDefault()
        }
      },
    },
  }
</script>
<style lang="scss" scoped>
  .watch-card {
    cursor: pointer;
    width: 15em;
  }
</style>