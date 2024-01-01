import React, {useEffect, useState} from "react";
import {Avatar, Button, Col, Row, Typography} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import {LogoutOutlined} from "@ant-design/icons";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";

const {Text,} = Typography;

interface Props {
}

const UserInfo: React.FC<Props> = (props) => {
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const screens = useBreakpoint();
    const [userMetadata, setUserMetadata] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUserMetadata(userInfo);
            }
        })();

    }, [isAuthenticated, fetchUserInfo]);

    return ((
        <Row style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
        }}>
            {userMetadata && isAuthenticated ? (
                <Row style={{
                    background: "white",
                    // paddingLeft: "12px",
                    // paddingRight: "16px",
                    borderRadius: "48px",
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    minHeight: "4vh",
                }}>
                    <Col span={3} offset={1}>
                        <Avatar src={userMetadata.picture}
                                size={36}
                                style={{marginRight: "12px"}}/>
                    </Col>
                    <Col span={12} offset={1}>
                        <div style={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            height: "100%",
                        }}>
                            <Text style={{
                                overflow: "hidden",
                                whiteSpace: "nowrap",
                                textOverflow: "ellipsis",
                                width: "100%",
                            }}>
                                {userMetadata.name}
                            </Text>
                        </div>
                    </Col>
                    <Col span={7}>
                        <div style={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            height: "100%",
                        }}>
                            <Button onClick={() => signOut(`${window.location.origin}/`)}
                                    style={{borderRadius: "30%"}}><LogoutOutlined/>
                            </Button>
                        </div>
                    </Col>


                </Row>
            ) : (
                <Button
                    onClick={() => signIn(`${window.location.origin}/callback`)}>Sign In
                </Button>
            )
            }
        </Row>
    ))
        ;
};

export default UserInfo;
