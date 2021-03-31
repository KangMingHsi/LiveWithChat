<template>
  <div class="watch-single">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            {{ watch.name }}
          </h1>
          <h2 class="subtitle ">
            {{ watch.date }}
          </h2>
        </div>
      </div>
    </section>
    <section class="watch-content">
      <div class="container">
        <p class="is-size-4 description">{{ watch.description }}</p>
        <p class="is-size-5"><strong>Location:</strong> {{ watch.location }}</p>
        <p class="is-size-5"><strong>Category:</strong> {{ watch.category }}</p>
        <div class="watch-images columns is-multiline has-text-centered">
          <div v-for="image in watch.images" :key="image.id" class="column is-one-third">
            <img :src="image" :alt="watch.name"/>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
  export default {
    name: 'WatchSingle',
    data() {
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
    created() {
      const ID = Number(this.$route.params.id)
      let watch = this.watches.find(watch => watch.id === ID)
      this.watch = watch
    }
  }
</script>
<style lang="scss" scoped>
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