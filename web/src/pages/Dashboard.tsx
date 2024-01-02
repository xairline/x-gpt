import React, {useEffect, useState} from "react";
import {Col, FloatButton, Row, Tabs, Typography} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import MapArch from "../components/mapArch";
import {CoffeeOutlined, IdcardOutlined, SettingOutlined} from "@ant-design/icons";
import Analytics from "../components/Analytics";

const {Title} = Typography;

interface Props {

}


const Dashboard: React.FC<Props> = () => {
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const [userMetadata, setUserMetadata] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUserMetadata(userInfo);
            }
        })();

    }, [isAuthenticated, fetchUserInfo]);
    return (
        <div style={{padding: "8", height: "100%"}}>
            <Row style={{height: "100%"}}>
                <Col span={window.innerHeight > window.innerWidth ? 24 : 8}
                     style={{height: window.innerHeight > window.innerWidth ? "40%" : "100%"}}>
                    <Tabs
                        defaultActiveKey="1"
                        type="card"
                        size={"middle"}
                        items={[{
                            label: `Analytics`,
                            key: "analytics",
                            children: <Analytics/>,
                        }]}
                        style={
                            {
                                marginRight: window.innerHeight > window.innerWidth ? 0 : "24px",
                            }
                        }
                    />
                </Col>
                <Col span={window.innerHeight > window.innerWidth ? 24 : 16}
                     style={{height: window.innerHeight > window.innerWidth ? "60%" : "100%"}}>
                    <MapArch/>
                </Col>

            </Row>

            <FloatButton.Group
                shape="circle"
                type="default"
                trigger={"hover"}
                style={{right: 64}}
                icon={<SettingOutlined/>}
                tooltip={<div>Settings</div>}
            >
                <FloatButton
                    shape="circle"
                    type="default"
                    icon={<IdcardOutlined/>}
                    tooltip={<div>Account</div>}
                />
                <FloatButton
                    shape="circle"
                    type="default"
                    icon={<CoffeeOutlined/>}
                    tooltip={<div>ChatGPT</div>}
                />
            </FloatButton.Group>

        </div>
    );
};

export default Dashboard;
