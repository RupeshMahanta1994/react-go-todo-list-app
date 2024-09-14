
import { Button, Container, Stack } from '@chakra-ui/react'
import './App.css'
import Navbar from './components/Navbar'
import TodoForm from './components/TodoForm'


function App() {


  return (
    <>
      <Stack>
        <Navbar />
        <Container>
          <TodoForm />
        </Container>
      </Stack>
    </>
  )
}

export default App
