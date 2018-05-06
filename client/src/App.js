import React, { Component } from 'react';
import axios from 'axios';
import logo from './logo.png';
import './App.css';
import { Icon } from 'react-icons-kit'
import {database} from 'react-icons-kit/fa/database'
import {Dialog
	, InputGroup
	, Button
	, Card
	, Elevation

} from '@blueprintjs/core'


class App extends Component {
	constructor(props){
		super(props);
		this.state ={
			loading : true,
			cluster: {
				nodes :[]
			},
			status: {
				columns :[],
				data : []
			},
			newNodeName: ''
		}
	}

	getCluster(){
		axios.get('/api/cluster').then(resp=>{
			this.setState({cluster: resp.data.data})
		})
	}
	getStatus(){
		axios.get('/api/status').then(resp=>{
			this.setState({status: resp.data.data})
		})
	}
	componentDidMount(){
		this.getCluster();
		this.getStatus();
	}

	renderEmptyPage(){
		return(
			<div className='empty-page'>
				<div className='empty-page-icon'>
					<Icon icon={database} size={64}/>
				</div>
				<div className='empty-page-h'>No Nodes found.</div>
				<div className='empty-page-t' onClick={this.openModal.bind(this)}>Create a node to initialise cluster.</div>
			</div>
		)
	}

	renderModal(){
		return(
			<Dialog
                    icon="databse"
                    isOpen={this.state.modalIsOpen}
                    onClose={this.closeModal.bind(this)}
                    title="Add Node"
                >
				<div className='modal-content'>
					{this.state.cluster.nodes.length >= 3 ?
					<div>
						<p> Due to RAM limitations, you can add a maximum of 3 nodes now.</p> 
						<Button style={{float: 'right'}} onClick={this.closeModal.bind(this)}>Ok</Button>
					</div>:
					<div>
						<p> Add a node to the cluster</p> 
						<InputGroup  placeholder="Node name" onChange={this.handleNodeNameChange.bind(this)}/>
						<br/>
						<Button style={{float: 'right'}} onClick={this.addNode.bind(this)}>Add Node</Button>
					</div>
					}
				</div>
        </Dialog>
		)
	}

	openModal() {
		this.setState({modalIsOpen: true});
	}

	handleNodeNameChange(evt){
		this.setState({newNodeName: evt.target.value});

	}

	stopNode(id){
		axios.post('/api/node/stop', {id}).then(resp=>{
			this.setState({cluster: resp.data.data})
		})
	}
	startNode(id){
		axios.post('/api/node/start', {id}).then(resp=>{
			this.setState({cluster: resp.data.data})
		})
	}

	addNode(){
		axios.post('/api/node/add', {name: this.state.newNodeName}).then(resp=>{
			this.setState({cluster: resp.data.data, newNodeName: '', modalIsOpen: false})
		})
	}
	
	
	
	closeModal() {
		this.setState({modalIsOpen: false});
	}

	renderContent(){
		const {nodes} = this.state.cluster;
		return(
			<div className='page'>
				<div className='row'>
					<div className='cluster'>
						<h3> Nodes</h3>
						<div className='node-container'>
						{nodes.map(node=>(
							<div className='node' data-active={node.active}>
								<Card interactive={false} elevation={node.active? Elevation.TWO: Elevation.ZERO}>
									<div className='empty-page-icon'>
										<Icon icon={database} size={64}/>
									</div>
									<div className='node-name'>{node.name}</div>
									<div className='node-status'>{node.status}</div>
									<div className='node-ip'>{node.ip}:{node.port}</div>
									<div>
										{node.active ?
											<Button onClick={this.stopNode.bind(this,node.id)}>Stop</Button>:
											<Button onClick={this.startNode.bind(this,node.id)}>Start</Button>
										}
									</div>

								</Card>
							</div>

						))}
						<div className='add-node' onClick={this.openModal.bind(this)}>
							+ 
						</div>
						</div>
					</div>
					<div className='status'>
						<table border="1">
							<tr> 
								{this.state.status.columns.map((column, i)=>(
									<th key={i}>{column}</th>
								))}
							</tr>
							{this.state.status.data.map((row, i)=>(
								<tr key={i}>
									{this.state.status.columns.map((column, j)=>(
										<td key={i}>{row[column]}</td>
									))}
								</tr>
							))}

						</table>
					</div>
				</div>
				<div className='row'>
					query
				</div>
			</div>
		)
	}
	render(){
		const {nodes}	= this.state.cluster
		return(
			<div>
				<div className='top-bar'>
					<img src={logo} className='logo'/>
				</div>
				<div>
					{nodes.length == 0 ? 
						this.renderEmptyPage():
						this.renderContent()
					}
				</div>
				{this.renderModal()}
			</div>
		)
	}
}

export default App;
