<template>
  <div class="relative wave md:block overflow-hidden md:h-20">
    <template v-if="isAdmin && finish">
      <AsyncAdminHeader></AsyncAdminHeader>
    </template>
    <template v-if="!isAdmin && finish">
      <AsyncUserHeader></AsyncUserHeader>
    </template>
  </div>
</template>

<script setup lang="ts">
const AsyncAdminHeader = defineAsyncComponent({
  loader: () => import('../admin/AdminHeader.vue')
})
const AsyncUserHeader = defineAsyncComponent({
  loader: () => import('../user/UserHeader.vue')
})

const { isAdmin, finish } = useAdminStatus()
</script>

<style scoped>
.wave::before,
.wave::after {
  content: '';
  position: absolute;
  left: 75%;
  min-width: 300vw;
  min-height: 300vw;
  background-color: rgba(148, 233, 255, 0.2);
  animation-name: rotate;
  animation-iteration-count: infinite;
  animation-timing-function: cubic-bezier(0.26, 0.86, 0.81, 0.26);
  overflow: hidden;
}

.wave::before {
  bottom: 5rem;
  border-radius: 45%;
  animation-duration: 10s;
}

.wave::after {
  bottom: 2rem;
  border-radius: 47%;
  animation-duration: 15s;
}

.wave {
  background: radial-gradient(
    circle,
    rgba(148, 233, 255, 0.2) 0%,
    rgba(200, 236, 255, 0.3) 40%,
    rgba(255, 255, 255, 0.2) 100%
  );
}

[data-theme='dark'] .wave::before,
[data-theme='dark'] .wave::after {
  display: none;
}

/* [data-theme='dark'] .wave::before,
[data-theme='dark'] .wave::after {
  background: radial-gradient(
    circle,
    rgba(148, 233, 255, 0.1) 0%,
    rgba(200, 236, 255, 0.1) 40%,
    rgba(255, 255, 255, 0.1) 100%
  );
} */

@keyframes rotate {
  0% {
    transform: translate(-50%, 0) rotateZ(0deg);
  }
  50% {
    transform: translate(-60%, -2%) rotateZ(180deg);
  }
  100% {
    transform: translate(-50%, 0%) rotateZ(360deg);
  }
}
</style>
