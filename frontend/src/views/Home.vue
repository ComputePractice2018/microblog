<template>
  <div class="home">
    <microblog title="Микроблоги"/>
  </div>
</template>

<script>
// @ is an alias to /src
import microblog from '@/components/microblog.vue'

export default {
  name: 'home',
  components: {
    microblog
  }
}
</script>
