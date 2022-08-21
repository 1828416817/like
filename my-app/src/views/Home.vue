<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <Index></Index>
  </div>
</template>

<script>
// @ is an alias to /src
import Index from "../components/Index";

export default {
  name: 'Home',
  components: {
    Index,
  }
}
</script>
