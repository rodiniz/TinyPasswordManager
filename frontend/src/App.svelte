
<script lang="ts">
  import { onMount } from "svelte";
  import { AddPassword, ListPasswords,RemovePassword} from "../wailsjs/go/main/App";
  import { ClipboardSetText} from "../wailsjs/runtime/runtime";
  import { Button, Col, Container, Form, FormGroup, Icon, Input, Label, Row, Table, Tooltip } from 'sveltestrap';
  let userName:string;
  let key: string;
  let valor:string;
  let passwords=[];
  let showPass=[false,false,false,false,false,false,false,false,false,false];
  onMount(async function () {   
     List()
   
  });
  async function Add(){
    const  error= await AddPassword({ username:userName, key, value: valor})
    key=""
    valor=""
    if(error!==""){
      alert(error)
      return 
    }
    
    await List()   
  }
  async function List(){
    var result= await ListPasswords(1,10)
    passwords= result.Passwords     
  }
  async function Delete(id:number){
    await RemovePassword(id)
    await List()  
  }
  function Copy(text:string){
    ClipboardSetText(text)
  }
  function generate(charset:string,length:number):string{   
    let retVal=""    
    for (var i = 0, n = charset.length; i < length; ++i) {
        retVal += charset.charAt(Math.floor(Math.random() * n));
    }
    return retVal;      
  }
  function GenerateRandomPassword(){
    
    let lowercaseCharSet = "abcdefghijkmnopqrstuvwxyz"; 
    let uppercaseCharSet = "ABCDEFGHJKLMNPQRSTUVWXYZ";   
    let numberCharSet = "23456789";   
    const specialCharSet = "!@#$%^*";
 
    let retVal:string=
      generate(lowercaseCharSet,3)+
      generate(numberCharSet,2)+
      generate(uppercaseCharSet,3)+
      generate(specialCharSet,1)+
      generate(lowercaseCharSet,2);

      valor=retVal;    
  }
  function hidechars(val :string):string{
    let masked="";
    for(let i=0; i < val.length;i++){
      masked+="*";
    }
    return masked;
  }
  function showHidePass(index:number){  
    showPass[index]=!showPass[index];
  }
</script>
<svelte:head>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css">
 
</svelte:head>
<main>
  
 
  <Container class="pt-5">
    <Form>
    <FormGroup>
      <Row>
        <Col xs="2"><Label color="black" for="key">User Name</Label></Col>
        <Col xs="6">
          <Input autocomplete="off" required class="input-box" bind:value={userName} id="key" type="text"/> 
        
        </Col>
      </Row>
      <Row>
        <Col xs="2"><Label color="black" for="key">Site</Label></Col>
        <Col xs="6">
          <Input autocomplete="off" required class="input-box" bind:value={key} id="key" type="text"/> 
        
        </Col>
      </Row>
      <Row class="mt-2">
        <Col xs="2"><Label for="value">Password</Label></Col>
        <Col xs="6">
          <Input autocomplete="off" required class="input" bind:value={valor} id="value" type="password"/>
        </Col>
        <Col xs="4">
          <Button type="button" id="controlledBtn" on:click={GenerateRandomPassword}><Icon name="clipboard"></Icon></Button>
          <Tooltip placement="right" target="controlledBtn">Generate random password</Tooltip>
        </Col>
      </Row>  
      <Row class="mt-2 mb-2">
        <Col xs="3"></Col>
        <Col xs="3"><Button type="submit" block color="primary"  on:click={Add}>Add</Button ></Col>
      </Row>
    </FormGroup> 
  </Form>
  </Container>
  {#if passwords.length >0}
    <Table bordered striped >
      <thead>
        <tr>
          <th colspan="2">Site</th>
          <th colspan="2">UserName</th>
          <th colspan="2">Password</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        {#each passwords as password, index}
         <tr>         
          <td> {password.key}</td>
          <td> {password.userName}</td>
          <td><Button type="button" on:click={()=>Copy(password.userName)}> <Icon name="clipboard"></Icon></Button></td>
          <td>
            {#if showPass[index]}
              {password.value}
            {:else}
              {hidechars(password.value)}
            {/if}
          </td>
          <td>
             <Button type="button" on:click={()=>showHidePass(index)}><Icon name="eye"></Icon></Button>
             <Button type="button" on:click={()=>Copy(password.value)}><Icon name="clipboard"></Icon></Button>
            </td>    
          <td><Button type="button" on:click={() => Delete(password.ID)} ><Icon name="trash" /></Button></td>            
        </tr>		    
	      {/each}
       
    
      </tbody>
    </Table>
  {/if}
   
 
</main>


