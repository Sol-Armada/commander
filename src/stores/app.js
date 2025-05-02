import { defineStore } from 'pinia'
import { getMember } from '@/api'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
    const member = ref(localStorage.getItem("member") || {})

    async function getSelf() {
        let m = await getMember('@me')
        if (m) {
            member.value = m
        }
    }

    return {
        member,
        getSelf,
    }
})
