<template>
  <div class="events container">
    <h2 class="subtitle is-3">
    Check out our videos or streams
    </h2>
    <div class="columns is-multiline">
      <div v-for="watch in watches" :watch="watch" :key="watch.id" class="column is-one-quarter">
        <router-link :to="'/watch/' + watch.id">
          <WatchCard :watch="watch" />
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
  import WatchCard from '@/components/WatchCard';
  export default {
    name: 'WatchesList',
    components : {
      WatchCard
    },
    data () {
      return {
        watch: {},
        watches: [
          {
            id: 1,
            name: 'Charity Ball',
            category: 'Fundraising',
            description: 'Spend an elegant night of dinner and dancing with us as we raise money for our new rescue farm.',
            featuredImage: 'https://placekitten.com/500/500',
            images: [
              'https://placekitten.com/500/500',
              'https://placekitten.com/500/500',
              'https://placekitten.com/500/500',
            ],
            location: '1234 Fancy Ave',
            date: '12-25-2019',
            time: '11:30'
          },
          {
            id: 2,
            name: 'Rescue Center Goods Drive',
            category: 'Adoptions',
            description: 'Come to our donation drive to help us replenish our stock of pet food, toys, bedding, etc. We will have live bands, games, food trucks, and much more.',
            featuredImage: 'https://placekitten.com/500/500',
            images: [
              'https://placekitten.com/500/500'
            ],
            location: '1234 Dog Alley',
            date: '11-21-2019',
            time: '12:00'
          }
        ]
      }
    },
  }
</script>
<style lang="scss" scoped>
  .events {
    margin-top: 100px;
    text-align: center;
  }
</style>