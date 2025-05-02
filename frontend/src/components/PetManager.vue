<template>
    <div class="pet-manager">
        <h1>Pet Manager</h1>

        <div v-if="loading" class="loading">Загрузка...</div>

        <div v-else>
        <label>
            ASCII:
            <textarea
            ref="asciiRef"
            v-model="pet.ascii"
            class="auto-resize"
            @input="resize($event.target)"
            />
        </label>

        <label>
            Description:
            <textarea
            ref="descRef"
            v-model="pet.description"
            class="auto-resize"
            @input="resize($event.target)"
            />
        </label>

        <button class="primary" @click="save">Сохранить</button>
        </div>

        <div v-if="toast" class="toast">{{ toast }}</div>
    </div>
</template>
  
<script setup lang="ts">
import { fetchPet, Pet, updatePet } from '@/services/api';
import { nextTick, onMounted, ref } from 'vue';

const pet     = ref<Pet>({ ascii: '', description: '' });
const loading = ref(true);
const toast   = ref('');

const asciiRef = ref<HTMLTextAreaElement|null>(null);
const descRef  = ref<HTMLTextAreaElement|null>(null);

function resize(el: HTMLTextAreaElement) {
    if (!el) return;
    el.style.height = 'auto';
    el.style.height = Math.min(el.scrollHeight, 4000) + 'px';
}

onMounted(async () => {
    pet.value = await fetchPet();

    await nextTick();
    await nextTick();

    setTimeout(() => {
      if (asciiRef.value) resize(asciiRef.value);
      if (descRef.value) resize(descRef.value);
    }, 0);


    loading.value = false;
});

async function save() {
    await updatePet(pet.value);
    toast.value = 'Данные сохранены!';
    setTimeout(() => (toast.value = ''), 3000);
}
</script>
  
<style scoped>
  .pet-manager { position: relative; }
  .loading    { font-style: italic; margin-bottom: 1em; }
  label {
    display: block;
    margin-bottom: 1em;
    font-weight: bold;
  }
  .auto-resize {
    width: 100%;
    min-height: 60px;
    max-height: 4000px;
    overflow: hidden;
    resize: none;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-family: monospace;
    box-sizing: border-box;
  }
  button.primary {
    padding: 8px 16px;
    background-color: #409eff;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .toast {
    position: absolute;
    top: -40px;
    right: 0;
    background: #48c774;
    color: white;
    padding: 8px 12px;
    border-radius: 4px;
    box-shadow: 0 2px 6px rgba(0,0,0,0.2);
  }
</style>
  