import { useEffect, useState } from "react";
import { Button, Text, TextInput, View } from "react-native";
import { EditModal } from "./modals/EditModal";

export default function Index() {
  const [dictionary] = useState<{[key : string] : string[]}>({});
  const [category, setCategory] = useState<string>('')
  const [visibility, setVisility] = useState<boolean>(false)

  const addCategory = () => {
    if(category != null && !dictionary[category]){
      dictionary[category] = []
    }
    else{
      alert('This category has been already created')
    }
    console.log(dictionary)

  }

  const cleanCategory = () => {
    setCategory('')
  }

  const showCategories = () =>{

    let categories: any[] = []

    for(let key in dictionary){
      categories.push(<Text>{key}<Button title="edit" onPress={()=>{if(!visibility){setVisility(true)}else{setVisility(false)}}}/></Text>)
    }
    return(
      <Text>{categories}</Text>
    )
  }



  return (
    <>
    <View
      style={{
        flex: 2,
        justifyContent: "center",
        alignItems: "center",
      }}
    >

      <TextInput 
      style={{padding: 5, borderColor: '#000000', borderWidth: 2}}
      placeholder="Digite a categoria da tarefa"
      value={category}
      onChangeText={setCategory}/>

      <Button 
      title="Click"
      onPress={()=>{addCategory(), cleanCategory()}}/>

      <View
          
        style={{
          backgroundColor: 'blue'
        }}>

        {showCategories()}

      </View>

      <EditModal visibility={visibility}/>

    </View>
    </>
    
  );
}
