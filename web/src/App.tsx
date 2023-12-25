import React, {useEffect, useState} from 'react';
import 'antd/dist/reset.css';
// import "@ant-design/plots/dist/index.css";
import './App.css';
import {Button, Col, ConfigProvider, Divider, Image, Layout, Menu, MenuProps, Row, Typography} from 'antd';
import UserInfo from "./components/UserInfo";
import {Link, Route, Routes} from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import {Header} from "antd/es/layout/layout";
import type {SelectProps} from 'antd/es/select';
import {AlertOutlined, DashboardOutlined, PhoneOutlined, SettingOutlined} from "@ant-design/icons";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";
import {useLogto, UserInfoResponse} from '@logto/react';
import Callback from "./pages/Callback";

const {Title} = Typography;
const {Footer, Content, Sider} = Layout;
type MenuItem = Required<MenuProps>['items'][number];
const largeStyle = {
    fontSize: '22px',
    padding: '8px 10px'
};
const items: MenuProps['items'] = [
    {
        key: 'dashboard',
        label: <><Link to={"dashboard"}>Dashboard</Link></>,
        icon: <DashboardOutlined style={largeStyle}/>,
        style: largeStyle
    },
    // {
    //     key: 'alerts',
    //     label: <><Link to={"alerts"}>Alerts</Link></>,
    //     icon: <AlertOutlined style={largeStyle}/>,
    //     style: largeStyle
    // },
    {
        key: "settings",
        label: <><Link to={"settings"}>Settings</Link></>,
        icon: <SettingOutlined style={largeStyle}/>,
        style: largeStyle
    },
    {
        key: "support",
        label: <><Link to={"support"}>Support & FAQ</Link></>,
        icon: <PhoneOutlined style={largeStyle}/>,
        style: largeStyle
    },

]

const App: React.FC = () => {
    let location = window.location.pathname;
    const screens = useBreakpoint();
    const [options, setOptions] = useState<SelectProps<object>['options']>([]);
    const {isAuthenticated, signIn, signOut, fetchUserInfo} = useLogto();
    const [user, setUser] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUser(userInfo);
            }
        })();
    }, [fetchUserInfo, isAuthenticated]);
    const [current, setCurrent] = useState(
        location === "/" || location === ""
            ? "dashboard"
            : location.split("/")[1]
    );

    return ((
        <ConfigProvider
            theme={{
                token: {
                    colorPrimary: '#00b96b',
                },
            }}
        >
            <Layout className="layout app">
                {!screens.sm ? <></> : <Header style={{
                    alignItems: 'center',
                    display: "flex",
                    height: !screens.xxl ? "10vh" : "12vh",
                }}>
                    <Col span={2}>
                        <div className="demo-logo">
                            <Image src={"logo512.png"} style={{maxHeight: "6vh"}}/>
                        </div>
                    </Col>
                    <Col span={20}><Menu
                        theme="dark"
                        mode="horizontal"
                        defaultSelectedKeys={['2']}
                        items={items}
                        style={{flex: 1, minWidth: 0, marginLeft: "12px"}}
                    /></Col>
                    <Col span={2}>{!isAuthenticated ?
                        <Button onClick={() => signIn(`${window.location.origin}/callback`)}>Sign In
                        </Button> :
                        (<><UserInfo/><Button onClick={() => signOut(`${window.location.origin}/callback`)}>Sign Out
                        </Button></>)}</Col>
                </Header>}

                <Layout>
                    {/*<Sider theme={'light'} style={{paddingLeft: "4px"}} collapsible>*/}

                    {/*    <Menu*/}
                    {/*        mode="inline"*/}
                    {/*        defaultSelectedKeys={[current]}*/}
                    {/*        style={{paddingTop: "24px"}}*/}
                    {/*        items={items}*/}
                    {/*    />*/}
                    {/*</Sider>*/}
                    <Layout style={{padding: '24px 24px 24px'}}>
                        <Content
                            style={{
                                padding: 24,
                                minHeight: 280,
                                background: "white",
                                overflow: "auto",
                            }}
                        >
                            <Routes>
                                <Route path={"/"} element={<Dashboard/>}/>
                                <Route path={"/dashboard"} element={<Dashboard/>}/>
                                <Route path="/callback" element={<Callback/>}/>
                            </Routes>
                        </Content>
                    </Layout>
                </Layout>
                {!screens.sm ? <Footer style={{
                    alignItems: 'center',
                }}>
                    <Row style={{width: "100%", height: "100%"}}>
                        <Col span={8}>
                        </Col>
                        <Col span={8}>
                            {!isAuthenticated ?
                                <Button onClick={() => signIn(`${window.location.origin}/callback`)}>Sign In
                                    ...</Button> :
                                <UserInfo/>}
                        </Col>
                        <Col span={8}>

                        </Col>
                    </Row>
                </Footer> : <></>}
            </Layout>
        </ConfigProvider>
    ));
};

export default App;
