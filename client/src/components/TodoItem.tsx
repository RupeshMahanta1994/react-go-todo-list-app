import { Badge, Center, Flex, Text } from '@chakra-ui/react'
import React from 'react'
import { IoCheckmarkDoneCircle } from "react-icons/io5";
import { MdDelete } from "react-icons/md";
import { Todo } from './TodoList';

const TodoItem = ({todo}:{todo:Todo}) => {
  return (
    <Flex align="center" justify="space-between">
      <Text>{todo.body}</Text>
      {
        todo.completed && <Badge colorScheme='green'>Done</Badge>
      }
      {
        !todo.completed && <Badge colorScheme='yellow'>In Progress</Badge>
      }
      <Flex>
      <IoCheckmarkDoneCircle onClick={()=>alert("I have been clicked")}/>
      <MdDelete />
      </Flex>

    </Flex>
  )
}

export default TodoItem