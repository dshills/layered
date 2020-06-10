import * as React from 'react';
// import {withRouter} from 'react-router-dom';
import './users.scss';
import * as moment from "moment";
import UserCard from "./user-card/userCard";
import PopUpModal from "../shared/popup-modal/popUpModal";
import UserForm from "./create-user/createUser";
import {NONNATEC_DATE_FORMAT} from "../../constants";


class Users extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            openModal: false
        }
    }

    addUser = () => {
        this.setState({
            openModal: true
        })
    };

    render() {
        return (
            <div className='users'>
                <div className='users-header'>
                    <div className='row'>
                        <div className='col-md-2'>
                            <div className='row'>
                                <div className='col-md-3'>
                                    <img src="/images/users@1.5x.svg" className='icon' alt=''/>
                                </div>
                                <div className='col-md-9'>
                                    <div className='date'>{moment().format(NONNATEC_DATE_FORMAT)}</div>
                                    <div className='name'>Users</div>
                                </div>
                            </div>
                        </div>
                        <div className='col-md-4 offset-md-4'>
                            <div className=" search">
                                <div className="input-group mb-4  input-icons">
                                    <i className="icon">
                                        <img src="/images/image-search.png" alt=""/>
                                    </i>
                                    <input className="search-box" placeholder="Search patient" type="text"/>
                                </div>
                            </div>
                        </div>
                        <div className='col-md-2'>
                            <button className='add-user' onClick={() => {
                                this.addUser()
                            }}>
                                <i><img src='/images/image-user-add@1.5x.svg' alt=""/></i>
                                <span className='text'>Create new user</span>
                            </button>
                        </div>
                    </div>
                    <hr/>
                    <div className='row'>
                        <div className='col-md-3 text-center '>
                            <div className='bottom-border'>
                                <UserCard/>
                            </div>
                        </div>
                        <div className='col-md-3 text-center '>
                            <div className='bottom-border'>
                                <UserCard/>
                            </div>
                        </div>
                        <div className='col-md-3 text-center '>
                            <div className='bottom-border'>
                                <UserCard/>
                            </div>
                        </div>
                        <div className='col-md-3 text-center '>
                            <div className='bottom-border'>
                                <UserCard/>
                            </div>
                        </div>
                    </div>
                </div>
                {
                    this.state.openModal &&
                    <PopUpModal
                        show={this.state.openModal}
                        onClose={() => {
                            this.setState({openModal: false})
                        }}
                        title={'Create New User'}>
                        <UserForm/>
                    </PopUpModal>
                }

            </div>
        );
    }
}

export default Users;