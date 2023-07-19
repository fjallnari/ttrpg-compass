<script lang="ts">
  import type TTRPGSystem from "../interfaces/TTRPGSystem";


  export let system: TTRPGSystem;

</script>

<div class="flex justify-center items-center w-full h-full">
  <div class="relative">
    <img src="../src/assets/base_star.svg" class="" alt="star">
    <!-- NORTH/CENTER TOP -->
    <div class="absolute left-[50%] translate-x-[-50%] bottom-[80%]">
      {#if system.Complexity > 1}
        <img src="../src/assets/star_{system.Complexity}.svg" class="" alt="star">
      {/if}
    </div>
    <!-- SOUTH/CENTER BOTTOM -->
    <div class="absolute left-[50%] translate-x-[-50%] top-[80%] rotate-180">
      {#if system.Progression > 1}
        <img src="../src/assets/star_{system.Progression}.svg" class="" alt="star">
      {/if}
    </div>
    <!-- EAST/CENTER RIGHT -->
    {#if system.Versatility > 1}
      <img src="../src/assets/star_{system.Versatility}.svg" class="absolute top-[50%] translate-y-[-50%] mx-[12px] rotate-90" alt="star">
    {/if}
    <!-- WEST/CENTER LEFT -->
    <!-- <div class="absolute left-[50%] top-[50%] translate-x-[-220%] translate-y-[-50%] rotate-[-90deg]">
      {#if system.Balance > 1}
        <img src="../src/assets/star_{system.Balance}.svg" class="" alt="star">
      {/if}
    </div> -->
  </div>
</div>