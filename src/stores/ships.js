// Ships
import { getShips } from '@/api'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useShipsStore = defineStore('ships', () => {
    const ships = ref([])
    const loading = ref(true)

    function getShip(id) {
        return ships.value.find(ship => ship.id === id)
    }

    if (ships.value.length === 0) {
        loading.value = true
        getShips()
            .then(data => {
                ships.value = data
                loading.value = false
            })
            .catch(error => {
                console.error('Error fetching ships:', error)
                loading.value = false
            })
    }

    return {
        ships,
        loading,
        getShip
    }
})
