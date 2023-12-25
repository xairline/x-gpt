import React, {useEffect, useState} from "react";
import {Avatar, Col, Divider, Dropdown, MenuProps, Row, Space, Spin, Typography} from "antd";
import { useLogto } from '@logto/react';
import {Link} from "react-router-dom";

const {Title, Paragraph} = Typography;

interface Props {
}

const UserInfo: React.FC<Props> = (props) => {
  const { isAuthenticated,fetchUserInfo } = useLogto();

  const [userMetadata, setUserMetadata] = useState();
  useEffect(() => {
    const getUserMetadata = async () => {
      const domain = window.location.origin + "/api";

      try {

        const metadataResponse = await fetchUserInfo();


        console.log(metadataResponse)
      } catch (e: any) {
        console.log(e.message);
      }
    };

  }, []);
  const items: MenuProps['items'] = [
    {
      key: '1',
      label: (
        <Link to={"settings"}>Settings</Link>
      ),
    },
    {
      key: '2',
      label: (
        <a onClick={() => {
          // logout({logoutParams: {returnTo: window.location.origin}})
        }}>
          Logout
        </a>
      ),
      disabled: false,
    },

  ];

  return ((
    <div style={{padding: "12px"}}>
      {userMetadata && isAuthenticated ? (
        <>
          <Dropdown menu={{items}} trigger={['click']}>
            <a onClick={(e) => e.preventDefault()}>
              <Avatar src={(userMetadata as any).picture} size={48}
                      style={{marginRight: "12px"}}/>
            </a>
          </Dropdown>

        </>
      ) : (
        <Spin size="large"/>
      )}
    </div>
  ));
};

export default UserInfo;
