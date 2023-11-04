/*
 * @Author: Xu Ning
 * @Date: 2023-11-04 14:09:25
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 15:36:37
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\test\CommentCom.spec.mjs
 */
import { mount } from '@vue/test-utils';
import CommentCom from '@/components/comment/CommentCom.vue';

describe('CommentCom', () => {
  it('renders the comment content correctly', () => {
    const comment = {
      user: {
        name: 'John Doe',
        avatar: 'path/to/avatar.png',
      },
      create_date: '2023-10-31',
      content: 'This is a comment.',
    };

    const wrapper = mount(CommentCom, {
      props: {
        comment,
      },
    });

    // 验证评论内容是否正确渲染
    expect(wrapper.text()).toContain(comment.content);

    // 验证评论作者名字是否正确渲染
    expect(wrapper.text()).toContain(comment.user.name);

    // 验证评论作者头像是否正确渲染
    const avatar = wrapper.find('.comment-avatar img');
    expect(avatar.attributes('src')).toBe(comment.user.avatar);

    // 验证评论创建日期是否正确渲染
    expect(wrapper.text()).toContain(comment.create_date);
  });

  it('emits "delete-comment" event when delete button is clicked', async () => {
    const comment = {
      user: {
        id: 'user123',
      },
      comment_id: 'comment123',
    };

    const wrapper = mount(CommentCom, {
      props: {
        comment,
      },
    });

    // 模拟点击删除按钮
    await wrapper.find('.comment-delete-button').trigger('click');

    // 验证是否正确触发了 "delete-comment" 事件，并传递了评论ID作为参数
    expect(wrapper.emitted('delete-comment')).toBeTruthy();
    expect(wrapper.emitted('delete-comment')[0][0]).toBe(comment.comment_id);
  });
});