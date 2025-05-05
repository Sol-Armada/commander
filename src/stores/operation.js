// Ships
import { defineStore } from 'pinia'
import { useMembersStore } from '@/stores/members'
import { useWebsocketStore } from '@/stores/websocket'
import { ref } from 'vue'
import { getOperation as apiGetOperation, getOperations } from '@/api'
import { WebSocketChannels, WebSocketTypes } from '@/classes/ws'
import { generateId } from '@/utility'
import { useAppStore } from './app'

const membersStore = useMembersStore()
const wsStore = useWebsocketStore()
const appStore = useAppStore()

export const useOperationStore = defineStore('operations', () => {
    const loading = ref(true)
    const currentOperation = ref(null)
    const operations = ref([])

    async function setOperation(operationId) {
        if (currentOperation.value && currentOperation.value.id === operationId) {
            console.debug("Operation already set")
            return true
        }

        const operation = operations.value.find(o => o.id === operationId)
        if (operation) {
            console.debug("Operation found in local store")
            currentOperation.value = operation
            loading.value = false
            return true
        }

        console.debug("Operation not found in local store, fetching from API")
        try {
            await getOperation(operationId)
        } catch (err) {
            console.error("Error fetching operation:", err)
            loading.value = false
            return false
        }
        return true
    }

    function addOperation(operation) {
        if (operations.value.find(o => o.id === operation.id)) {
            return
        }
        operations.value.push(operation)
    }

    async function getOperation(opId) {
        let op = await apiGetOperation(opId)
        if (op) {
            currentOperation.value = op
            loading.value = false
        }
        loading.value = false
    }

    function getOnStandBy() {
        return currentOperation.value.onStandBy
    }

    function putOnStandby(memberId) {
        if (currentOperation.value.onStandBy === null || currentOperation.value.onStandBy === undefined) {
            currentOperation.value.onStandBy = []
        }
        if (currentOperation.value.onStandBy.find(m => m.id === memberId)) {
            return
        }

        removeMember(memberId)

        let member = membersStore.getMember(memberId)
        currentOperation.value.onStandBy?.push(member)

        save()
    }

    function removeMember(memberId) {
        currentOperation.value.onStandBy = currentOperation.value.onStandBy.filter(m => m.id !== memberId)

        currentOperation.value.ships.forEach(ship => {
            ship.positions.forEach(position => {
                if (position.member && position.member.id === memberId) {
                    position.member = null
                }
            })
        })

        currentOperation.value.squads.forEach(squad => {
            squad.members = squad.members.filter(m => m.id !== memberId)
        })

        save()
    }

    function renameOperation(newName) {
        currentOperation.value.name = newName
        save()
    }

    function renameSquad(squadId, newName) {
        const squad = currentOperation.value.squads.filter(s => s.id === squadId)[0]
        squad.name = newName
        save()
    }

    function addShip(ship) {
        console.debug("adding ship", ship)
        if (currentOperation.value.ships === undefined || currentOperation.value.ships === null) {
            currentOperation.value.ships = []
        }
        ship.id = generateId()
        currentOperation.value.ships.push(ship)
        save()
    }

    function addToShipPosition(shipId, positionId, memberId) {
        const ship = currentOperation.value.ships.filter(s => s.id === shipId)[0]
        if (ship === undefined) {
            console.error("Ship not found")
            return
        }
        const position = ship.positions.filter(position => position.id == positionId)[0]
        if (position === undefined) {
            console.error("Position not found")
            return
        }
        if (position.member) {
            console.error("Position already has a member")
            return
        }

        let member = membersStore.getMember(memberId)
        if (member === undefined) {
            console.error("Member not found")
            return
        }

        removeMember(memberId)
        position.member = member
        save()
    }

    // function clearPosition(shipId, positionId) {
    //     const ship = currentOperation.value.ships.filter(s => s.id === shipId)[0]
    //     const position = ship.positions.filter(position => position.id == positionId)[0]
    //     position.member = null
    // }

    function addMemberToSquad(squadId, memberId) {
        const squad = currentOperation.value.squads.filter(s => s.id === squadId)[0]

        removeMember(memberId)

        let member = membersStore.getMember(memberId)
        squad.members.push(member)
        save()
    }

    function removeShip(shipId) {
        currentOperation.value.ships = currentOperation.value.ships.filter(s => s.id !== shipId)
        save()
    }

    function addSquad() {
        if (currentOperation.value.squads === undefined || currentOperation.value.squads === null) {
            currentOperation.value.squads = []
        }

        const squad = {
            id: generateId(),
            name: 'Squad ' + (currentOperation.value.squads.length + 1),
            members: [],
        }

        currentOperation.value.squads.push(squad)
        save()
        return squad
    }

    function removeSquad(squadId) {
        const squad = currentOperation.value.squads.filter(s => s.id === squadId)[0]
        if (!squad) {
            return
        }

        currentOperation.value.squads = currentOperation.value.squads.filter(s => s.id !== squad.id)
        save()
    }

    function hasMember(memberId) {
        if (currentOperation.value.onStandBy?.find(m => m.id === memberId)) {
            return true
        }
        if (currentOperation.value.squads?.find(s => s.members.find(m => m.id === memberId))) {
            return true
        }
        if (currentOperation.value.ships?.find(s => s.positions.find(p => p.member && p.member.id === memberId))) {
            return true
        }

        return false
    }

    function save() {
        wsStore.send(WebSocketChannels.OPERATIONS, WebSocketTypes.UPDATE, currentOperation.value)
    }

    function isCreator() {
        if (currentOperation.value === null) {
            return false
        }
        if (currentOperation.value.creator === null || currentOperation.value.creator === undefined) {
            return false
        }
        if (appStore.member.id === currentOperation.value.creator) {
            return true
        }
        return false
    }

    getOperations().then(res => {
        console.debug("Operations fetched", res)
        operations.value = res
        loading.value = false
    }).catch(err => {
        console.error("Error fetching operations:", err)
        loading.value = false
    })

    wsStore.addChannel(WebSocketChannels.OPERATIONS, (res) => {
        if (res.type === WebSocketTypes.CREATED) {
            addOperation(res.data);
        }
        if (res.type === WebSocketTypes.UPDATED) {
            if (currentOperation.value !== null && res.data.id === currentOperation.value.id) {
                currentOperation.value = res.data;
            }

            const operation = operations.value.find(o => o.id === res.data.id)
            if (operation) {
                operations.value[operations.value.indexOf(operation)] = res.data
            } else {
                operations.value.push(res.data)
            }
        }
        if (res.type === WebSocketTypes.DELETED) {
            if (res.data.id === currentOperation.value.id) {
                currentOperation.value = null;
            }
        }
    })

    return {
        loading,
        setOperation,
        getOperation,
        addOperation,
        addShip,
        addToShipPosition,
        // clearPosition,
        putOnStandby,
        removeShip,
        removeMember,
        addSquad,
        removeSquad,
        renameOperation,
        renameSquad,
        addMemberToSquad,
        currentOperation,
        getOnStandBy,
        hasMember,
        operations,
        isCreator,
    }
})
