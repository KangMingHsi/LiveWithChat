<template>
  <v-container>
    <div class="video-section">
    </div>
    <v-row>
      <v-col>
        <p>TITLE</p>
        <p>xxxx view</p>
        <v-divider></v-divider>
        <p>dddddddddddddddddd</p>
      </v-col>

      <v-col>
        <v-card
          elevation="16"
          max-width="400"
          class="mx-auto"
          max-height="500"
        >
          <v-virtual-scroll
            :items="watches"
            :item-height="50"
            height="400"
          >
            <template v-slot:default="{ item }">
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>{{ item.Title }}</v-list-item-title>
                </v-list-item-content>
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
              >
                <v-icon>mdi-send</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>

    <div class="watch-single">
      <section class="hero is-primary">
        <div class="hero-body">
          <div class="container">
            <h1 class="title">
              {{ watch.Title }}
            </h1>
            <h2 class="subtitle ">
              {{ watch.Description }}
            </h2>
          </div>
        </div>
      </section>
      <section class="watch-content">
        <div class="container">
          <p class="is-size-4 description">{{ watch.Description }}</p>
          <p class="is-size-5"><strong>Location:</strong> {{ watch.Description }}</p>
          <p class="is-size-5"><strong>Category:</strong> {{ watch.Description }}</p>
          <div class="watch-images columns is-multiline has-text-centered">
            <div v-for="image in watch.images" :key="image.id" class="column is-one-third">
              <img :src="image" :alt="watch.name"/>
            </div>
          </div>
        </div>
      </section>
    </div>
  </v-container>
</template>
<script>
  export default {
    name: 'WatchSingle',
    data() {
      return {
        message: "",
        watch: {},
      }
    },
    created() {
      // const ID = Number(this.$route.params.id)
      // let watch = this.watches.find(watch => watch.ID === ID)
      // this.watch = watch
    }
  }
</script>
<style lang="scss" scoped>
  .video-section {
    width: 100%;
    height: 480px;
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