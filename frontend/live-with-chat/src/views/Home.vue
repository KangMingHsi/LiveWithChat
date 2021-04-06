<template>
<div class="home">
  <section class="hero is-dark">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          Welcome to the Live With Chat          
        </h1>
        <h2 class="subtitle">
          Enjoy videos and live streams here
        </h2>
        <v-dialog
         v-model="showRegister"
         width="500"
        >
          <template v-slot:activator="{ on, attrs }">
            <div class="button-block">
              <button class="button is-xl is-dark"
               v-bind="attrs"
               v-on="on"
              >
                Sign Up to Browse Subscribers
              </button>
            </div>
          </template>
          <RegisterDialog
            ref="registerDialog"
            @closeRegister="closeRegister"
          />
        </v-dialog>
      </div>
    </div>
  </section>
  <WatchesList />
</div>
</template>

<script>
import WatchesList from '@/components/WatchesList';
import RegisterDialog from '../components/RegisterDialog'

export default {
  name: 'home',
  components: {
    WatchesList,
    RegisterDialog
  },
  data() {
    return {
      showRegister: false
    }
  },
  watch: {
    showRegister: function(val) {
      val || this.$refs.registerDialog.clear()
    }
  },
  methods: {
    closeRegister() {
      this.showRegister = false
    }
  },
}
</script>

<style lang="scss" scoped>
  .hero {    
    text-align: center;
    background-image: url('../assets/galaxy.jpg');
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    height: 400px;
  }
  .hero-body .title {
    text-shadow: 4px 4px 4px rgba(0, 0, 0, 0.6);
    padding: 40px 0 20px 0;
    font-size: 60px !important;
  }
  .subtitle {
    text-shadow: 4px 4px 4px rgba(0, 0, 0, 0.7);
    font-size: 30px !important;
  }
  .button-block {
    text-align: center;
    margin-left: auto;
    margin-right: auto;
    width: 100%;
    position: absolute;
    bottom: -150px;
    .button {
      margin-right: 50px;
      padding-left: 50px;
      padding-right: 50px;
    }
    .welcome {
      width: 400px;
      padding: 10px;
      margin-left: auto;
      margin-right: auto;
    }
  }
  .is-xl {
    font-size: 1.7rem;
  }
</style>