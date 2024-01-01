import React, {useEffect, useState} from 'react';
import 'antd/dist/reset.css';
// import "@ant-design/plots/dist/index.css";
import './App.css';
import {Button, Col, ConfigProvider, Image, Layout, MenuProps, Row, Typography} from 'antd';
import UserInfo from "./components/UserInfo";
import {Route, Routes} from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import type {SelectProps} from 'antd/es/select';
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

const App: React.FC = () => {
    let location = window.location.pathname;
    const screens = useBreakpoint();
    const [options, setOptions] = useState<SelectProps<object>['options']>([]);
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const [user, setUser] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUser(userInfo);
            }
        })();
    }, [isAuthenticated, fetchUserInfo]);
    const [current, setCurrent] = useState(
        location === "/" || location === ""
            ? "dashboard"
            : location.split("/")[1]
    );

    return ((
        <ConfigProvider
            theme={{
                token: {
                    colorPrimary: '#006363'
                },
            }}
        >
            <Layout className="layout app">
                {!screens.sm ? <></> :
                    <Row style={{background: "#006363"}}>
                        <Col flex={"auto"}>
                            <Row style={{
                                display: "flex",
                                height: "100%",
                            }}>
                                <Image src={"logo512.png"}
                                       style={{maxHeight: "8vh", objectFit: "contain", margin: "12px 24px 12px"}}>
                                </Image>

                            </Row>
                        </Col>
                        <Col flex={"auto"}>
                            <Title level={1} style={{
                                color: "white", display: "flex",
                                justifyContent: "center",
                                alignItems: "center",
                                height: "100%",
                            }}>
                                X Airline
                            </Title>
                        </Col>
                        <Col flex={"auto"}>
                            <div style={{
                                display: "flex",
                                justifyContent: "flex-end",
                                alignItems: "center",
                                height: "100%",
                                marginRight: "24px",
                            }}>
                                <UserInfo/>
                            </div>
                        </Col>
                    </Row>
                }

                <Layout>
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
                {!screens.sm ? <Footer style={{
                    alignItems: 'center',
                }}>
                    <Row style={{width: "100%", height: "100%"}}>
                        <Col span={1}>
                        </Col>
                        <Col span={22}>
                            {!isAuthenticated ?
                                <Button onClick={() => signIn(`${window.location.origin}/callback`)}>Sign In
                                    ...</Button> :
                                <UserInfo/>}
                        </Col>
                        <Col span={1}>

                        </Col>
                    </Row>
                </Footer> : <></>}
            </Layout>
        </ConfigProvider>
    ));
};

export default App;
