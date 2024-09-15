import { Flex, Stack, Text } from '@chakra-ui/react'
import React, { useEffect, useState } from 'react'
import TodoItem from './TodoItem';
import { useQuery } from '@tanstack/react-query';


export type Todo={
    _id:number,
    body:string,
    completed:boolean
}

const TodoList = () => {
    const getTodos=async()=>{
        try{

            const res=await fetch("http://127.0.0.1:5001/api/todos")
            const data=res.json();
            console.log(data)
        }catch(error){
            console.log("Error in getting the todos",error)
        }
    }
    useEffect(()=>{
        getTodos()
    },[])
   //const [isLoading, setIsLoading] = useState(true);
    //Query functin
    // const {data:todos,isLoading}=useQuery<Todo[]>({
    //     queryKey:['todos'],
    //     queryFn:async ()=>{
    //         try{
    //             const response=await fetch("http://127.0.0.1:5000/api/todos");
    //             const data=await response.json();
    //             if(!response.ok){
    //                 throw new Error(data.error || "Something went wrong")
    //             }
    //             return data ||[];

    //         }catch(error){
    //             console.log("Error in fetching the todos",error)
    //         }
    //     }
    // })
    return (
        <>
            <Text>
                Today's Tasks
            </Text>
            <Stack>
                {todos?.map((todo)=>(
                    <TodoItem key={todo._id} todo={todo}/>
                ))}
            </Stack>

        </>
    )
}

export default TodoList

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