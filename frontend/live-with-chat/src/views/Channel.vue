<template>
  <v-row class="row">
    <div class="text-center"
    >

      <WatchesList
        class="watch-list"
        :conditions="query"
        :editable="true"
      />

      <UploadVideoDialog 
        class="upload"
      />
    </div>
  </v-row>
</template>

<script>
  import WatchesList from '@/components/WatchesList';
  import UploadVideoDialog from '@/components/UploadVideoDialog'

  export default {
    name: 'Channel',
    components: {
      WatchesList,
      UploadVideoDialog
    },
    data() {
      return {
        id: this.$route.params.id,
        query: {
          "user_id": this.id,
        },
      }
    },
    created() {
      // Get self video list
    },
    methods: {
      
    },
  }
</script>

<style lang="scss" scoped>

</style>