<template>
  <v-container>
    <v-card
      width="400px"
      height="300px"
    >
      <v-img
        height="200px"
        src="https://placekitten.com/400/400"
      >
      </v-img>
      <small class="watch-duration">{{ toTimeString() }}</small>
      <v-row
        align="start"
        dense
        style="padding-top: 5px; padding-left: 5px;"
      >
        <v-col
          cols="2"
        >
          <v-avatar
            size="28px"
            color="primary"
          >
            <v-icon dark>
              mdi-account-circle
            </v-icon>
          </v-avatar>
        </v-col>
        <v-col cols="10">
          <v-row
            no-gutters
            align="center"
          >
            <div class="ellipsis">
              <h2 class="is-size-4 has-text-weight-bold">{{ watch.Title }}</h2>
            </div>
          </v-row>
          <v-row
            no-gutters
            align="center"
          >
            <div class="ellipsis">
              <span>{{ watch.OwnerID }}</span>
            </div>
          </v-row>
        </v-col>
      </v-row>
      
    </v-card>
  </v-container>
</template>
<script>
  export default {
    props: [
      'watch'
    ],
    methods: {
      toTimeString() {
        var seconds = Math.round(this.watch.Duration / 1000000000)
        var hours = parseInt( seconds / 3600 );
        seconds = seconds % 3600;

        var minutes = parseInt( seconds / 60 );
        seconds = seconds % 60;

        minutes = (minutes < 10) ? "0" + minutes : minutes;
        seconds = (seconds < 10) ? "0" + seconds : seconds;
        var hms = minutes+":"+seconds;
        if (hours > 0) {
          hours = (hours < 10) ? "0" + hours : hours;
          hms = hours+":"+hms
        }
        return hms
      },
    },
  }
</script>
<style lang="scss" scoped>
  .ellipsis {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    h2 {
      left: 10px;
      font-size: .85em !important;
      text-align: left;
      text-overflow: ellipsis;
      line-height: normal;
    }
    span {
      font-size: .8em;
      text-align: left;
      // width: 100%;
      // bottom: 10px;
      text-overflow: ellipsis;
      line-height: normal;
      left: 10px;
    }
  }
  .watch-duration {
    background-color: #151515;
    color: #FFF;
    font-size: .75em;
    padding: 2px 10px;
    position: absolute;
    top: 0;
    right: 0;
  }
</style>