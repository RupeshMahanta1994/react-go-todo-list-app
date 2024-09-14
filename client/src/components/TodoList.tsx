import { Text } from '@chakra-ui/react'
import React, { useState } from 'react'
const todos = [
    {
        _id: 1,
        body: "Buy this",
        completed: true
    },
    {
        _id: 2,
        body: "Buy that",
        completed: false
    }
]
const TodoList = () => {
    const [isLoading, setIsLoading] = useState(true);
    return (
        <>
            <Text>
                Today's Tasks
            </Text>
            {
                isLoading ? && todo
   }

        </>
    )
}

export default TodoList