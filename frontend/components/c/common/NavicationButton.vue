<!-- components/NavigationButton.vue -->
<template>
  <v-btn :value="value" :to="to">
    <span :class="{ 'primary--text': isActive }">{{ label }}</span>
    <v-icon :color="isActive ? 'primary' : ''">{{ icon }}</v-icon>
  </v-btn>
</template>

<script>
export default {
  name: 'NavigationButton',
  props: {
    value: {
      type: String,
      required: true,
    },
    to: {
      type: String,
      required: true,
    },
    icon: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
  },
  computed: {
    isActive() {
      return this.$route.path === this.to
    },
  },
}
</script>
