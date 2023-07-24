<script lang="ts">
    import Icon from "@iconify/svelte";

    export let color: string = "[#E5E7EB]";

</script>

<div class="text-{color} text-2xl opacity-80">
    <div class="absolute left-[50%] translate-x-[-50%] bottom-[82%]">
        <Icon icon="mdi:brain" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[50%] translate-x-[-50%] bottom-0">
        <Icon icon="mdi:sword-cross" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[27%] translate-x-[-50%] bottom-6">
        <Icon icon="mdi:ruler-square-compass" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[74%] translate-x-[-50%] bottom-6">
        <Icon icon="mdi:scroll-text" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[19%] translate-x-[-50%] bottom-[41%]">
        <Icon icon="mdi:puzzle" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[81%] translate-x-[-50%] bottom-[41%]">
        <Icon icon="mdi:scale-balance" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[74%] translate-x-[-50%] bottom-[74%]">
        <Icon icon="mdi:stairs" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
    <div class="absolute left-[27%] translate-x-[-50%] bottom-[74%]">
        <Icon icon="mdi:shape" class="hover:text-goldenrod transition-colors cursor-pointer"/>
    </div>
</div>