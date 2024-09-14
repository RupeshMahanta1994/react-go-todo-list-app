import { Alert, Button, Flex, Input, Spinner } from "@chakra-ui/react"

import React, { useState } from "react"
import { IoMdAdd } from "react-icons/io"


const TodoForm = () => {
    const [newTodo, setNewTodo] = useState("")
    const [isPending, setIsPending] = useState(false)
    const createTodo = async (e: React.FormEvent) => {
        e.preventDefault();
        alert("Todo created")
    }
    return (
        <form onSubmit={createTodo}>
            <Flex gap={5}>
                <Input type="text" onChange={(e) => setNewTodo(e.target.value)}
                    ref={input => input && input.focus()}
                />
                <Button type="submit">
                    {isPending?<Spinner size={"xs"}/>:<IoMdAdd size={30}/>}
                </Button>
            </Flex>
        </form>
    )
}

export default TodoForm